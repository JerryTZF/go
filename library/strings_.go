/**
 * Created by GoLand
 * Time: 2021/9/24 10:24 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: strings_.go
 * Desc: strings包常规操作
 */
package library

import (
	"bytes"
	"fmt"
	"strings"
)

// StrIndex 位置相关
func StrIndex() {
	str1 := "hello world"

	// 检测字符串是否包含子串 bool
	fmt.Println(strings.Contains(str1, "orl"))

	// 检测字符串是否包含子串中的任一个字符
	fmt.Println(strings.ContainsAny(str1, "abcd"))

	// 字符串中包含子串的个数
	fmt.Println(strings.Count(str1, "l"))

	// 字符换是否含有指定的后缀
	fmt.Println(strings.HasSuffix(str1, "ld"))

	// 字符串是否含有指定的前缀
	fmt.Println(strings.HasPrefix(str1, "he"))

	// 字符串A在字符串B中出现的位置
	fmt.Println(strings.Index(str1, " "))

	// 字符A在字符串B中出现的位置
	fmt.Println(strings.IndexByte(str1, 'o'))

	// 字符串substr在str中出现的最后一次的位置
	fmt.Println(strings.LastIndex(str1, "ll"))

	// 字符c在str中出现的最后一次的位置
	fmt.Println(strings.LastIndexByte(str1, 'l'))

	// 返回满足fn=true的字符串的索引值,类似PHP的array_filter()
	// index = 2
	index := strings.IndexFunc("aLlow", func(r rune) bool {
		if r == 'l' {
			return true
		} else {
			return false
		}
	})
	fmt.Println(index)
}

// StrSplit 字符串切割相关
func StrSplit() {
	str := "hello world !"

	// 以空白字符分割字符串，并且填充到一个切片内 []string
	fmt.Println(strings.Fields(str)[1])

	// 将字符串str以满足f(r) == true的字符分割，返回一个切片
	sl1 := strings.FieldsFunc("hello world!", func(r rune) bool {
		if r == 'o' {
			return true
		} else {
			return false
		}
	})
	fmt.Println(sl1)

	// 将字符串str以sep字符串分割，返回一个切片
	fmt.Println(strings.Split("aaa-bbb-ccc", "-"))

	// 将字符串str以sep字符串分割,分割后最后附上sep，返回一个切片
	fmt.Println(strings.SplitAfter("aaa-bbb-ccc", "-"))

	// 将字符串str以sep字符串分割,分割后最后附上sep，返回一个切片,决定返回的切片数
	fmt.Println(strings.SplitAfterN("aaa-bbb-ccc", "-", 2))

	// 将字符串str以sep字符串分割，返回一个切片,决定返回的切片数
	fmt.Println(strings.SplitN("aaa-bbb-ccc", "-", 2))
}

// StrUpper 大小写转换
func StrUpper() {
	// 每个单词首字母大写
	fmt.Println(strings.Title("hello world!"))

	// 全部转为大写
	fmt.Println(strings.ToUpper("hello world!"))

	// 全部转为小写
	fmt.Println(strings.ToLower("hello World!"))
}

// StrTrim 修剪字符串
func StrTrim() {
	// 将字符串str首尾包含在cutset中的任一字符去掉返回
	fmt.Println(strings.Trim("xCxabcxxx", "x"))

	// 去除字符串两端的空格符
	fmt.Println(strings.TrimSpace(" xxx    "))
}

// StrCompare 字符串比较
func StrCompare() {
	// 字符串是否相等(按照字典序比较字符串大小)
	if strings.Compare("aaa", "AAA") == 0 {
		fmt.Println("字符串相等")
	} else {
		fmt.Println("字符串不相等")
	}

	// 将字符串重复n次
	fmt.Println(strings.Repeat("k", 10))

	// 替换字符串
	fmt.Println(strings.Replace("hello world!", "l", "L", 1))

	// 将切片中的字符串，拼接成一个新的字符串，并且以-作为分隔符
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
}

func Commonly() {
	// 字符串拼接
	fmt.Println("hello" + " world!")
	// 字符串拼接
	fmt.Println(string(append([]byte("hello"), " world!"...)))
	// 字符串拼接
	buf := new(bytes.Buffer)
	buf.WriteString("hello")
	buf.WriteString(" world!")
	fmt.Println(buf.String())
	// 拼接字符串
	builder := new(strings.Builder)
	builder.WriteString("hello")
	builder.WriteString(" world!")
	fmt.Println(builder.String())
	//拼接字符串
	fmt.Println(strings.Join([]string{"hello", " world!"}, ""))
}
