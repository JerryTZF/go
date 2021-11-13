## 一些关键点

- struct可以理解为轻量的(没有成员方法)类
- 一个类型的成员方法是不允许重载的
- 类型的方法必须要在同一个包内
- 对象的字段（属性）不应该由 2 个或 2 个以上的不同线程在同一时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用包 sync中的方法

---
## 关于内嵌

> 1、将一个类型A放入另一个类型(例入struct类型)B中，那么类似B继承了A
> 2、类：提供一组 作用于 共同类型的方法集(所谓的隐式定义)

> 下面模拟一个多重继承
> ```go
> package pkg
> import "fmt"
> 
> // ************
> // 演示多继承
> // ************
> type Employee struct {
>   Title  string
>   Salary int
> }
>
> type Son struct {
>    Mother string
>    Father string
> }
>    
> type Jerry struct {
>    Ability string
>    Hobby   string
> }
>
> type Me struct {
>   *Employee
>   *Jerry
>   Son
> }
> 
> func (j *Jerry) ChangeHobby(h string) {
>   j.Hobby = h
> }
>
> func DemoForExtend() {
>   e := &Employee{Title: "技术组长", Salary: 15000}
>   j := &Jerry{Ability: "PHP", Hobby: "乒乓球"}
>   s := Son{Mother: "王", Father: "田"} 
>   me := &Me{e, j, s}
>
>   // 一定注意:这里不要并发修改hobby
>   me.ChangeHobby("网球")
>   j.ChangeHobby("不是网球")
>   fmt.Println(me.Hobby)
> }
>   
> ```
