package main

import "fmt"

type User struct {

}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

func a()(resutl int)  {
	if true {
		resutl = 1
	}
	return resutl
}

func main() {
	c:=a()
	fmt.Printf("%d",c)
	/*type myint1 int
	type myint2 = int
	var i int = 9
	var i1 myint1 = myint1(i)
	var i2 myint2 = i
	fmt.Println(i1, i2)*/
	/*考点：**Go 1.9 新特性 Type Alias **
	基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。 MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转； MyInt2称之为alias，可以直接赋值。*/
/*	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()*/
}
