/**
 * Created by GoLand
 * Time: 2021/11/24 11:02 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: rpc_server.go
 * Desc:
 */
package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct{}

func (m *MathUtil) CalculatedCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}

func main() {
	err := rpc.Register(new(MathUtil))
	if err != nil {
		panic(err.Error())
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}

	_ = http.Serve(listen, nil)
}
