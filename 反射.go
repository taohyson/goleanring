package main

import (
	"fmt"
	"reflect"
)

func main(){
	var s int  =12
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
}