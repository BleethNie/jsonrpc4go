package main

import (
	"github.com/BleethNie/jsonrpc4go"
)

type IntRpc struct{}

type Params struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	C int `json:"c"`
}

func (i *IntRpc) Add(params *Params, result *int) error {
	a := params.A + params.B
	*result = any(a).(int)
	return nil
}

func (i *IntRpc) Add2(params *Params, result *[]Result) error {
	*result = append(*result, Result{C: params.A})
	*result = append(*result, Result{C: params.B})
	return nil
}

func main() {
	s, _ := jsonrpc4go.NewServer("tcp", 3232)

	s.RegisterWithName(new(IntRpc), "")
	s.Start()
}
