/**
 * Created by GoLand
 * Time: 2021/10/7 23:30
 * Author: JerryTian<tzfforyou@163.com>
 * File: json_.go
 * Desc: json 相关操作
 */
package library

import (
	"encoding/json"
	"fmt"
)

// 定义好json对应的结构体
type Address struct {
	Type    string
	City    string
	Country string
}

type Card struct {
	firstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func JsonDemo() {
	a1 := &Address{
		Type:    "private",
		City:    "Beijing",
		Country: "China",
	}

	a2 := &Address{
		Type:    "public",
		City:    "HeBi",
		Country: "China",
	}
	card := Card{
		firstName: "Tian",
		LastName:  "ZhaoFan",
		Address:   []*Address{a1, a2},
		Remark:    "none",
	}

	// 结构体转Json(只有公共字段才会被解析为json)
	if bytes, err := json.Marshal(card); err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	// 更加复杂的结构体(带有slice、map)
	s := struct {
		Name  string
		Age   int
		Roles []string
		Skill map[string]float64
		Info  []*Address
	}{
		Name:  "Jerry",
		Age:   27,
		Roles: []string{"Master", "Worker"},
		Skill: map[string]float64{"python": 90.0, "PHP": 95.5, "Golang": 92.5},
		Info:  []*Address{a1, a2},
	}

	if bytes, err := json.Marshal(s); err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
	}
}
