package main

import (
	"fmt"
)

type Rectangle struct{
	W, H float64
}

//kế thừa phương thức
type ColoredRectangle struct{
	Rectangle

	Color string
}

func (r Rectangle) Area() float64{
	return r.W * r.H
}

func main(){
	rect1 := Rectangle{2, 3}

	fmt.Println(rect1.Area())

	var rect2 ColoredRectangle
	rect2.W = 3
	rect2.H = 5
	fmt.Println(rect2.Area())
}