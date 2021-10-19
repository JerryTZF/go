/**
 * Created by GoLand
 * Time: 2021/10/18 3:27 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: reflect_.go
 * Desc:
 */
package library

import (
	"fmt"
	"reflect"
)

func ReflectDemo() {
	x := 3.14
	y := "Pi"
	fmt.Printf("x type is :%v; y type is :%v\n", reflect.TypeOf(x), reflect.TypeOf(y))
	fmt.Printf("x value is :%v; y value is :%v\n", reflect.ValueOf(x), reflect.ValueOf(y))

	// 类型判断
	if reflect.TypeOf(x).Kind() == reflect.Float64 {
		fmt.Println("x type is", reflect.Float64)
	}

	// 通过 reflect.ValueOf 的 Elem() 可以得到传入的变量的指针指向的具体对象
	address := new(Address)
	address.Type = "public"
	address.Country = "China"
	address.City = "ShenZhen"

	v := reflect.ValueOf(address).Elem()

	// 通过反射修改变量
	if reflect.ValueOf(address).Elem().CanSet() {
		v.Set(reflect.ValueOf(&Address{
			Type:    "private",
			Country: "China",
			City:    "HeBi",
		}).Elem())
	}
	fmt.Println(v)
	fmt.Println(address)

	// 通过反射修改切片指定位置的值
	intS := []int{128, 256, 512}
	intV := reflect.ValueOf(intS)
	e := intV.Index(0)
	if e.CanSet() {
		e.SetInt(64)
		fmt.Println(intS)
	}

	// 通过反射修改整个切片
	intE := reflect.ValueOf(&intS).Elem()
	if intE.CanSet() {
		newS := []int{1, 2, 3}
		intE.Set(reflect.ValueOf(newS))
		fmt.Println(intS)
	}

	// 结构体反射
	addressIns := Address{"protect", "zhanJiang", "China"}

	structType := reflect.TypeOf(addressIns)

	fmt.Println(structType.NumField()) // 返回结构体成员数量

	for i := 0; i < structType.NumField(); i++ {
		// 结构体字段类型
		filedType := structType.Field(i)
		fmt.Println(filedType)
	}

	// 获取结构体中的tag
	stuIns := Stu{"Zhaofan Tian", 27, "PingPang", 54.5}
	stuType := reflect.TypeOf(stuIns)
	if fieldName, ok := stuType.FieldByName("Hobby"); ok {
		fmt.Println(fieldName.Tag.Get("json"))
	}

	// 反射获取结构体字段值
	stuValue := reflect.ValueOf(stuIns)
	fmt.Println(stuValue.FieldByName("Age").Int())
	fmt.Println(stuValue.FieldByName("Name").String())

	// 反射修改结构体字段值
	stuNameValue := reflect.ValueOf(&stuIns.Name)
	stuNameValue.Elem().SetString("ZHAOFAN TIAN")
	fmt.Println(stuIns)

	// 通过反射调用结构体方法
	f := stuValue.MethodByName("GetName")
	res := f.Call([]reflect.Value{reflect.ValueOf("Jerry")})
	fmt.Println(res)

}
