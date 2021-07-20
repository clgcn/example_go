package main

// 接口，我们定义了一个接口，任何类型只要实现了这个接口定义的方法，就是实现了这个接口，也可以说这个类型也是属于这个接口类型
// 如果我们在struct中定义了一个interface类型，那么表示这个property可以接受任何type，只要这个type实现了这个接口
// for range 还可以用在chan上，用来循环chan接收到的元素
// func main() {
// links := []string{
// 	"http://google.com",
// 	"http://baidu.com",
// 	"http://stackoverflow.com",
// 	"http://golang.org",
// 	"http://amazon.com",
// }

// c := make(chan string)

// for _, link := range links {
// 	go checkLink(link, c)
// }

// // //这个版本每次main goroutine都会把l传递给child goroutine，所以可以一直循环
// for l := range c {
// 	go func(link string) {
// 		time.Sleep(5 * time.Second)
// 		checkLink(link, c)
// 	}(l)
// }

// //这个版本并没有参数传递给func literal，也就是chile goroutine所以l的值在上面的for循环结束之后就停在了最后的值
// //那么引用的就是最后的值，也就是main goroutine上的值
// for l := range c {
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		checkLink(l, c)
// 	}()
// }

// }

// func checkLink(link string, c chan string) {
// 	time.Sleep(5 * time.Second)
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		c <- link
// 		return
// 	}
// 	fmt.Println(link, "is up!")
// 	c <- link
// }

// type Person interface {
// 	cook()
// }

// type woman struct {
// 	name string
// 	age  int
// }

// type man struct {
// 	name string
// 	age  int
// }

// func main() {
// 	// zcj := woman{
// 	// 	name: "zhouchengjiao",
// 	// 	age:  10,
// 	// }
// 	// print(zcj)

// 	gcl := man{
// 		name: "gcl",
// 		age:  11,
// 	}
// 	print(gcl)
// 	gcl.cook()
// }

// func (w woman) cook() {
// 	fmt.Println("Woman Cook!")
// }

// func (m man) cook() {
// 	fmt.Println("man cook!")
// }

// func (m man) wash() {
// 	fmt.Println("man wash!")
// }

// func print(p Person) {
// 	if _, ok := p.(interface{ wash() }); ok {
// 		fmt.Println("implement the wash method")
// 	}
// 	p.cook()
// }

type scene struct {
	name string
	age  int
}

// type Game struct {
// 	scenes map[string]*scene
// }

// func main() {
// 	s := scene{
// 		name: "zhangsan",
// 		age:  10,
// 	}
// 	sm := make(map[string]*scene)
// 	sm["zs"] = &s
// 	// 下面这话的意思就是说如果map的key存在，就ok啦，如果不存在就不ok啦
// 	if s, ok := sm["zs"]; !ok {
// 		panic("hehe")
// 	} else {
// 		fmt.Println(s.age)
// 	}
// 	// fmt.Println(g.scenes["zs"].name)
// }
