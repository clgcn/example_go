package main

import "fmt"

// make 只用于slice map channel的初始化，返回的是类型为slice，map，或者channel的值
// new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针，并且内存空间会清零
func main() {
	b := new(int)
	fmt.Println(*b)

	c := make([]int, 0)
	c = append(c, 1, 2, 3, 4, 5)
	fmt.Println(c)
}
