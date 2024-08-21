package client

import (
	"bytes"
	"fmt"
	"github.com/BleethNie/jsonrpc4go/common"
	"github.com/BleethNie/jsonrpc4go/discovery"
	"net"
	"strconv"
	"time"
)

type Tcp struct {
	Name      string
	Protocol  string
	Address   string
	Discovery discovery.Driver
}

type TcpClient struct {
	Name        string
	Protocol    string
	Address     string
	Discovery   discovery.Driver
	RequestList []*common.SingleRequest
	Options     TcpOptions
	Pool        *Pool
}

type TcpOptions struct {
	ReqAddEof     string //请求时添加eof字符
	RespCheckEof  string //自定义结束符
	RespMaxLength int64
	RespAddSuffix string //有时结束符和实际数据混合,需要把数据补上
}

func (p *Tcp) NewClient() Client {
	return NewTcpClient(p.Name, p.Protocol, p.Address, p.Discovery)
}

func NewTcpClient(name string, protocol string, address string, dc discovery.Driver) *TcpClient {
	options := TcpOptions{
		"",
		"\n",
		512 * 1024,
		"",
	}
	pool := NewPool(name, address, dc, PoolOptions{5, 5})
	return &TcpClient{
		name,
		protocol,
		address,
		dc,
		nil,
		options,
		pool,
	}
}

func (c *TcpClient) BatchAppend(method string, params any, result any, isNotify bool) *error {
	singleRequest := &common.SingleRequest{
		method,
		params,
		result,
		new(error),
		isNotify,
	}
	c.RequestList = append(c.RequestList, singleRequest)
	return singleRequest.Error
}

func (c *TcpClient) BatchCall() error {
	var (
		err error
		br  []any
	)
	for _, v := range c.RequestList {
		var (
			req any
		)
		realMethod := fmt.Sprintf("%s/%s", c.Name, v.Method)
		if c.Name == "" {
			realMethod = fmt.Sprintf("%s", v.Method)
		}
		if v.IsNotify == true {
			req = common.Rs(nil, realMethod, v.Params)
		} else {
			req = common.Rs(strconv.FormatInt(time.Now().Unix(), 10), realMethod, v.Params)
		}
		br = append(br, req)
	}
	bReq := common.JsonBatchRs(br)
	bReq = append(bReq)
	err = c.handleFunc(bReq, c.RequestList)
	c.RequestList = make([]*common.SingleRequest, 0)
	return err
}

func (c *TcpClient) SetOptions(tcpOptions any) {
	c.Options = tcpOptions.(TcpOptions)
}

func (c *TcpClient) SetPoolOptions(poolOption any) {
	c.Pool.SetOptions(poolOption.(PoolOptions))
}

func (c *TcpClient) Call(method string, params any, result any, isNotify bool) error {
	var (
		err error
		req []byte
	)
	realMethod := fmt.Sprintf("%s/%s", c.Name, method)
	if c.Name == "" {
		realMethod = fmt.Sprintf("%s", method)
	}
	if isNotify {
		req = common.JsonRs(nil, realMethod, params)
	} else {
		req = common.JsonRs(strconv.FormatInt(time.Now().Unix(), 10), realMethod, params)
	}
	req = append(req, []byte(c.Options.ReqAddEof)...)
	err = c.handleFunc(req, result)
	return err
}

func (c *TcpClient) handleFunc(req []byte, result any) error {
	var (
		err  error
		conn net.Conn
	)

	conn, err = c.Pool.Borrow()
	if err == nil {
		_, err = conn.Write(req)
	} else {
		conn, err = c.Pool.BorrowAfterRemove(conn)
		if err != nil {
			c.Pool.Remove(conn)
			return err
		}
		_, err = conn.Write(req)
		if err != nil {
			c.Pool.Remove(conn)
			return err
		}
	}

	defer c.Pool.Release(conn)

	eofb := []byte(c.Options.RespCheckEof)
	eofl := len(eofb)
	var (
		data []byte
	)
	l := 0
	for {
		var buf = make([]byte, c.Options.RespMaxLength)
		n, err := conn.Read(buf)
		if err != nil {
			if n == 0 {
				return err
			}
			common.Debug(err.Error())
		}
		l += n
		data = append(data, buf[:n]...)
		if bytes.Equal(data[l-eofl:], eofb) {
			break
		}
	}
	//移除EOF
	data = data[:l-eofl]
	//添加
	pSuffix := []byte(c.Options.RespAddSuffix)
	if len(pSuffix) > 0 {
		data = append(data, pSuffix...)
	}
	err = common.GetResult(data, result)
	return err
}
