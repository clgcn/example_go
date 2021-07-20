package main

import "fmt"

func main() {
	// fmt.Println(rt())
	d()
	// b()
	// fmt.Println(c())
	// fmt.Println(rt())

}

func rt() []int {
	fmt.Println("Lock")
	// defer fmt.Println("Unlock")
	products := make([]int, 0)
	products = append(products, 1)
	fmt.Println("Unlock")
	return products
}

// 当defer被声明时，参数会实时解析，所以这里得到的结果是0
func a() int {
	i := 0
	defer fmt.Println(i)
	return i + 1
}

// golang按照先进后出的规则依次调用defer，所以这里的结果是3210
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// return 之后进行函数调用，这里先return 1，然后调用literal
// return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func d() (a int) {
	fmt.Println("start")
	defer func() { a++; fmt.Printf("unlock %d\n", a) }()
	fmt.Println(a)
	fmt.Println("end")
	return a
}
