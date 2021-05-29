package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	// cấu trúc của string là reflect.StringHeader
	var s = "hello world"

	fmt.Println(s[:5])
	fmt.Println(s[6:])

	// string is immutable (bất biến) or read-only
	//s[0] = 'H'
	//fmt.Println(s)

	fmt.Printf("length of s: %d\n", len(s))
	fmt.Printf("length of s: %d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
}