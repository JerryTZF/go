## 一些关键点

- struct可以理解为轻量的(没有成员方法)类
- 一个类型的成员方法是不允许重载的
- 类型的方法必须要在同一个包内
- 对象的字段（属性）不应该由 2 个或 2 个以上的不同线程在同一时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用包 sync中的方法

---
## 关于内嵌

> 1、将一个类型A放入另一个类型(例入struct类型)B中，那么类似B继承了A

> 下面模拟一个多重继承
> ```go
> package pkg
> import "fmt"
> 
> type A struct {
>   name string
> }
> 
> type B struct{
>   brade string
> }
> 
> type C struct {
>   A
>   *B
> }
> 
> func Demo()  {
>   a := &C{A{"Jerry"},&B{"Apple"}}
>   fmt.Println(a.brade)
> }
> ```
