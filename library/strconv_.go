/**
 * Created by GoLand
 * Time: 2021/9/25 2:44 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: strconv_.go
 * Desc: strconv包常规操作(字符串和基础数据类型转换)
 */
package main

import (
	"fmt"
	"strconv"
)

// 语法糖转换
func CandyChange() {
	// 将整数转换为十进制字符串形式
	fmt.Printf("%T;%v\n", strconv.Itoa(100), 100)
	fmt.Printf("%T;%v\n", string(100), 100)

	// 十进制字符串转整型
	n, _ := strconv.Atoi("16")
	fmt.Printf("%T;%v\n", n, n)
}

func FormatToString() {
	// 将整数转换为字符串形式。base 表示转换进制，取值在 2 到 36 之间。
	// 将整数转换为十进制字符串形式
	fmt.Printf("%T;%v\n", strconv.FormatInt(100, 10), 100)

	// 布尔值转字符串
	fmt.Println(strconv.FormatBool(false))

	// 将浮点数转换为字符串类型
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// 具体参数含义参考源码
	fmt.Println(strconv.FormatFloat(3.141592654, 'f', 5, 64))
}

func ParseStringToOther() {
	// 尝试将字符串转布尔值
	if b, err := strconv.ParseBool("K"); err == nil && b {
		fmt.Println("字符串转布尔成功，且布尔值为true")
	} else {
		// 它接受假值：0, f, F, FALSE, false, False
		// 它接受真值：1, t, T, TRUE, true, True;
		fmt.Println(b, err)
	}

	// 字符串解析为整型
	// base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
	// bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
	if n, err := strconv.ParseInt("10", 10, 64); err == nil {
		fmt.Printf("%T;%v\n", n, n)
	} else {
		fmt.Println(n, err)
	}

	// 将一个表示浮点数的字符串转换为 float 类型
	if f, err := strconv.ParseFloat("3.141592654", 64); err == nil {
		fmt.Printf("%T;%v\n", f, f)
	} else {
		fmt.Println(f, err)
	}
}

func FormatToStringAppend() {
	var sl []byte
	sl = strconv.AppendInt(sl, 42, 10)
	fmt.Println(sl)
}

func EscapeString() {
	s := "hello/t世界! \n"
	// 去除转义
	s = strconv.Quote(s)
	fmt.Println(s)

	// 非 ASCII 字符和不可打印字符会被转义
	s = strconv.QuoteToASCII(s)
	fmt.Println(s)
}
