package main

import (
	"fmt"
	"unsafe"
)

type Rect struct {
	Width  int
	Height int
}

func main() {
	var r = Rect{50, 50}
	var width = *(*int)(unsafe.Pointer(&r))
	var height = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&r)) + uintptr(8)))
	fmt.Println(width)
	fmt.Println(height)

	var s = []int{1,2,3,4,5,6,7,8,9,10}
	var address = (**[10]int)(unsafe.Pointer(&s))
	var len = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	var cap = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(address, *len, *cap)
	var body = **address
	for i:=0; i< len(body); i++ {
		fmt.Printf("%d ", body[i])
	}
}
