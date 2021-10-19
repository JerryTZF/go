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

func Struct2Json() {
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

// 通过tag重新命名
type Stu struct {
	Name   string  `json:"name"`            // 重命名Name -> name
	Age    int     `json:"-"`               // 显式忽略该字段
	Hobby  string  `json:"hobby,omitempty"` // 该字段有值的时候输出,反之(零值)的时候不输出
	Weight float64 `json:"weight,string"`   // 最终输出值的类型与定义不符需要转换时
}

func (s Stu) GetName(str string) string {
	return s.Name + "_" + str
}

func RenameField() {
	s := Stu{"Jerry", 12, "", 50.25}
	if bytes, err := json.Marshal(s); err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes)) // {"name":"Jerry"}
	}
}

// json转struct
type Opus struct {
	Date  string
	Title string
}

type Actress struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       []Opus
}

func Json2Struct() {
	jsonStr := `{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":[
         {
            "date":"2013",
            "title":"《阿娜尔罕》"
         },
         {
            "date":"2014",
            "title":"《逆光之恋》"
         },
         {
            "date":"2015",
            "title":"《克拉恋人》"
         }
      ]
   }`
	jsonData := []byte(jsonStr)
	s := new(Actress)
	if err := json.Unmarshal(jsonData, s); err == nil {
		fmt.Println(s)
	}
}
