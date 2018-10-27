package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("first")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()
	f()
}

func f() {
	fmt.Println("test")
	panic(1)
	fmt.Println("test2")
}
