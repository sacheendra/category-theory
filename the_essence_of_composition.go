package main

import "fmt"

func main() {
	fmt.Println(compose(id, id)("Hello, playground"))
}

func id(a interface{}) interface{} { return a }

func compose(f1 func(interface{}) interface{}, f2 func(interface{}) interface{}) func(interface{}) interface{} {
	return func(a interface{}) interface{} {
		return f2(f1(a))
	}
}
