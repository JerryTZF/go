/**
 * Created by GoLand
 * Time: 2021/10/13 10:57 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: error_.go
 * Desc:
 */
package pkg

import (
	"fmt"
	"strconv"
)

func ErrorReturned(f float64) (float64, error) {
	if f <= 0 {
		//return 0, errors.New("Args lte 0")
		return 0, fmt.Errorf("Args is : %.4f", f)
	}
	return strconv.ParseFloat(fmt.Sprintf("%.4f", f*f), 64)
}

// 自定义错误(实现Go异常接口)
type FileError struct {
	Op   string
	Name string
	Path string
}

func newFileError(op, name, path string) *FileError {
	return &FileError{
		Op:   op,
		Name: name,
		Path: path,
	}
}

func (f *FileError) Error() string {
	return fmt.Sprintf("路径：%s; 文件：%s; 行为：%s", f.Path, f.Name, f.Op)
}

func NewFileError() error {
	return newFileError("读取文件", "README.md", "/opt/www/")
}

// 系统抛出异常的时候，Defer仍然可以运行
func DeferError() int {
	defer func() {
		fmt.Println("=====")
	}()
	panic("手动报错")
	a := 10 / 1
	return a
}

// defer恢复异常
func DeferError_(a, b int) (r int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			r = -1
		}
	}()
	r = a / b
	return r
}
