/**
 * Created by GoLand
 * Time: 2021/10/26 2:45 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: http_.go
 * Desc: net/http 标准库使用
 */
package library

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SayHello struct {
	Name string
}

// 实现Handler接口
func (s *SayHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!!!" + s.Name))
}

func Run() {
	http.Handle("/hello", &SayHello{"XiaoMing"})
	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!!!"))
	})

	// DefaultServeMux 是 ServeMux(多路复用器变量)的一个变量
	// ServeMux 是 Handler(interface) 的实现
	e := http.ListenAndServe("127.0.0.1:8080", nil)
	if e != nil {
		panic(e)
	}
}

// 自定义多路复用器
func Run_() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!!!"))
	})

	s := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func Run__() {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Welcome!\n")
	})

	router.GET("/hello/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "hello, %s!\n", params.ByName("name"))
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Get() {
	url := "https://service.cqysda.com/admin/captcha.json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("unique", "asd")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	response, err := http.DefaultClient.Do(req)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := io.ReadAll(response.Body)
		_ = ioutil.WriteFile("./captcha.png", body, 0666)
	} else {
		fmt.Println(response.Header)
	}
}

func UrlEncodePost() {
	url_ := "https://service.cqysda.com"
	response, err := http.Post(url_, "application/x-www-form-urlencoded", strings.NewReader("a=b"))
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode == 200 {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(response.Header)
	}
}

func FormPost() {
	url_ := "https://service.cqysda.com"
	payload := url.Values{"key": {"val2"}, "key_bak": {"bak1"}}
	response, err := http.PostForm(url_, payload)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode == 200 {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(response.Header)
	}
}

func JsonPost() {
	url_ := "https://service.cqysda.com"
	payload := strings.NewReader("{\"a\":{\"b\":[\"b1\",\"b2\",\"b3\"]},\"c\":{\"d\":\"e\"}}")
	req, _ := http.NewRequest("POST", url_, payload)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode == 200 {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(response.Header)
	}
}
