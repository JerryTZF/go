/**
 * Created by GoLand
 * Time: 2021/10/26 2:45 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: http_.go
 * Desc: net/http 标准库使用
 */
package library

import "net/http"

type SayHello struct {
	Name string
}

// 实现Handler接口
func (s *SayHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!!!" + s.Name))
}

func Run() {
	http.Handle("/hello", &SayHello{"XiaoMing"})
	e := http.ListenAndServe("127.0.0.1:8080", nil)
	if e != nil {
		panic(e)
	}
}
