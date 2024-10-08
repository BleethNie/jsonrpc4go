package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/BleethNie/jsonrpc4go/common"
	"github.com/BleethNie/jsonrpc4go/discovery"
	"golang.org/x/time/rate"
	"log"
	"net"
	"sync"
	"time"
)

type Tcp struct {
	Hostname string
	Port     int
}

type TcpServer struct {
	Hostname  string
	Port      int
	Server    common.Server
	Options   TcpOptions
	Event     chan int
	Discovery discovery.Driver
}

type TcpOptions struct {
	ReqCheckEof  string
	RespAddEof   string
	ReqMaxLength int64
}

func (p *Tcp) NewServer() Server {
	options := TcpOptions{
		"\n",
		"",
		1024 * 1024 * 2,
	}
	return &TcpServer{
		p.Hostname,
		p.Port,
		common.Server{
			sync.Map{},
			common.Hooks{},
			nil,
		},
		options,
		make(chan int, 1),
		nil,
	}
}

func (s *TcpServer) Start() error {
	// Register services
	if s.Discovery != nil {
		register := func(key, value interface{}) bool {
			go s.DiscoveryRegister(key, value)
			return true
		}
		s.Server.Sm.Range(register)
	}
	// Start the server
	var addr = fmt.Sprintf("%s:%d", s.Hostname, s.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	log.Printf("Listening tcp://%s:%d", s.Hostname, s.Port)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Event <- 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		go s.handleFunc(ctx, conn)
	}
}

func (s *TcpServer) DiscoveryRegister(key, value interface{}) bool {
	err := s.Discovery.Register(key.(string), "tcp", s.Hostname, s.Port)
	if err == nil {
		return true
	}
	time.Sleep(REGISTRY_RETRY_INTERVAL * time.Millisecond)
	s.DiscoveryRegister(key, value)
	return false
}

func (s *TcpServer) Register(m any) {
	s.Server.Register(m)
}

func (s *TcpServer) RegisterWithName(m any, name string) {
	s.Server.RegisterWithName(m, name)
}

func (s *TcpServer) SetOptions(tcpOptions any) {
	s.Options = tcpOptions.(TcpOptions)
}

func (s *TcpServer) SetDiscovery(d discovery.Driver, hostname string) {
	s.Discovery = d
	s.Hostname = hostname
	var err error
	if s.Hostname == "" {
		s.Hostname, err = GetHostname()
		if err != nil {
			common.Debug(err.Error())
		}
	}
}

func (s *TcpServer) SetRateLimit(r rate.Limit, b int) {
	s.Server.RateLimiter = rate.NewLimiter(r, b)
}

func (s *TcpServer) SetBeforeFunc(beforeFunc func(id any, method string, params any) error) {
	s.Server.Hooks.BeforeFunc = beforeFunc
}

func (s *TcpServer) SetAfterFunc(afterFunc func(id any, method string, result any) error) {
	s.Server.Hooks.AfterFunc = afterFunc
}

func (s *TcpServer) GetEvent() <-chan int {
	return s.Event
}

func (s *TcpServer) handleFunc(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	select {
	case <-ctx.Done():
		return
	default:
		//	do nothing
	}
	eofb := []byte(s.Options.ReqCheckEof)
	eofl := len(eofb)
	for {
		var (
			data []byte
		)
		l := 0
		for {
			var buf = make([]byte, s.Options.ReqMaxLength)
			n, err := conn.Read(buf)
			if err != nil {
				if n == 0 {
					return
				}
				common.Debug(err.Error())
			}
			l += n
			data = append(data, buf[:n]...)
			if bytes.Equal(data[l-eofl:], eofb) {
				break
			}
		}
		res := s.Server.Handler(data[:l-eofl])
		res = append(res, eofb...)
		conn.Write(res)
	}
}
