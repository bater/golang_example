package main

import "fmt"

func main() {

	var i *int // 宣告i是一個int的指標，目前還不知道會指向哪邊

	a := 10 // a佔用了一個記憶體空間

	i = &a // 將i指到a的記憶體位置

	fmt.Println(i)  // i所指到的記憶體位置
	fmt.Println(*i) // *代表顯示該記憶體位置的值

	// 宣告一個空的int
	n := new(int)

	// 直接把值指向記憶體位置
	*n = 2

	fmt.Println(n)
	fmt.Println(*n)

	foo_value(*n)
	foo_point(n)
}

func foo_value(x int) {
	fmt.Println(&x) // function內x的記憶體位置
}

func foo_point(x *int) {
	fmt.Println(x) // function內x的記憶體位置
}
