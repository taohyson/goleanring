package main

import "fmt"

func printa(content string) {
	fmt.Println("print " + content)
}

type A func(string) //回调函数
func main() {

	var v interface{}
	var a A = printa
	v = a
	(v.(A))("fdsfd")

	var typea interface{} = "fdsfds"
	switch v := typea.(type) {
	case string:
		fmt.Println("typea.(type):", "string", v)
	}
}
