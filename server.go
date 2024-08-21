package jsonrpc4go

import (
	"errors"
	"github.com/BleethNie/jsonrpc4go/server"
)

func NewServer(protocol string, port int) (server.Server, error) {
	var p server.Protocol
	switch protocol {
	case "http":
		p = &server.Http{"", port}
	case "tcp":
		p = &server.Tcp{"", port}
	default:
		return nil, errors.New("The protocol can not be supported")
	}
	return server.NewServer(p), nil
}

// 指定hostname参数
func NewHServer(protocol string, hostname string, port int) (server.Server, error) {
	var p server.Protocol
	switch protocol {
	case "http":
		p = &server.Http{hostname, port}
	case "tcp":
		p = &server.Tcp{hostname, port}
	default:
		return nil, errors.New("The protocol can not be supported")
	}
	return server.NewServer(p), nil
}
