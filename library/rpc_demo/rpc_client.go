/**
 * Created by GoLand
 * Time: 2021/11/24 11:12 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: rpc_client.go
 * Desc:
 */

package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	// 定义float32类型的指针
	//var resp *float32
	// 远程过程调用(同步)
	//err = client.Call("MathUtil.CalculatedCircleArea", 3.0, &resp)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(*resp)

	// 远程过程调用(异步)
	var respSync *float32
	syncCall := client.Go("MathUtil.CalculatedCircleArea", 3.0, &respSync, nil)
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(*respSync)

}
