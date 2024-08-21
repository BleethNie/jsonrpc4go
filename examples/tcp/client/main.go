package main

import (
	"fmt"
	"github.com/BleethNie/jsonrpc4go"
	"github.com/BleethNie/jsonrpc4go/client"
)

type Params struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	C int `json:"c"`
}

func main() {
	c, _ := jsonrpc4go.NewClient("", "tcp", "127.0.0.1:3232")                                                            // or "IntRpc/Add2", "int_rpc.Add2", "IntRpc.Add2"
	c.SetOptions(client.TcpOptions{ReqAddEof: "\n", RespMaxLength: 1024 * 512, RespAddSuffix: "}", RespCheckEof: "}\n"}) // Custom package EOF when the protocol is tcp

	// notify
	result2 := []Result{}
	err2 := c.Call("Add2", Params{1, 6}, &result2, true)
	// data sent: {"jsonrpc":"2.0","method":"IntRpc/Add2","params":{"a":1,"b":6}}
	// data received: {"jsonrpc":"2.0","result":{"c":7}}
	fmt.Println(err2)    // nil
	fmt.Println(result2) // {7}

}
