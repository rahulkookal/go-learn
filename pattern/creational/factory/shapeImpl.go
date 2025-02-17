package factory

import "fmt"

type Circle struct {
}

func (c Circle) Draw() {
	fmt.Println("Circle")
}

type Rectangle struct {
}

func (r Rectangle) Draw() {
	fmt.Println("Rectangle")
}
