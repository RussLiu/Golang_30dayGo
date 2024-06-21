package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup
var money int = 1500
var mu sync.Mutex

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	fmt.Println("Hello World!")
	call1("Russ", "Night")

	fmt.Println(call2("Russ", "Morning"))

	//迴圈
	fmt.Println(plusminas(3, 2))

	for i := 1; i <= 5; i++ {
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print("*")

		}
		fmt.Println("")
	}

	//function多值
	fmt.Println(total(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//funciton變數宣告
	foo2 := func() {
		fmt.Println("Hi " + name)
	}
	foo2()

	//function回傳function(用法不明)
	bar := foo1()
	fmt.Println(bar())

	//Callback
	compile([]string{"Winform", "ASP.NET", "Golang"}, func(n string) {
		fmt.Println(n)
	})

	//defer
	project()

	//傳指標(pass-by-pointer)
	fmt.Println("name舊記憶體位址", &name)
	change(&name)
	fmt.Println("更改name記憶體位址內容")
	fmt.Println("name新記憶體位址", &name)
	fmt.Println("name新內容", name)

	//傳struct
	p := stuff{"pancil", 10}
	fmt.Println(p.name, "價格為:", p.price)
	inprice(&p)
	fmt.Println(p.name, "改價格為:", p.price)

	//自宣告自用
	func() {
		fmt.Println("lalala")
	}()

	//Pointer
	var o *int
	a := 10
	o = &a

	fmt.Println(o)
	fmt.Println(*o)

	c := 10
	fmt.Println("main記憶體位址", &c)
	foo3(c)
	foo4(&c)

	//Array
	var x [10]int

	fmt.Println("Array Length:", len(x))
	for i, n := range x {
		fmt.Println("index:", i, ";value:", n)
	}

	//Map
	//h:=map[string]int{}
	//var h =make(map[string]int)
	//h:=make(map[string]int)

	h := map[string]int{"Tony": 168, "Mary": 159, "Alen": 185}
	fmt.Println(h["Tony"])

	h["George"] = 163
	fmt.Println(h["George"])
	fmt.Println("Map Length:", len(h))

	if key, exists := h["Tony"]; exists {
		fmt.Println("Find ", key)
	}

	//Slice List
	s := make([]string, 0, 5)
	s = append(s, "element1", "element2", "element3", "element4")
	fmt.Println("Slice Length:", len(s))
	fmt.Println("Slice Capacity:", cap(s))

	for _, n := range s {
		fmt.Println(n)
	}

	//Strucct
	//t:=person{"Tony",169}
	t := person{name: "Tony", height: 169}
	fmt.Println("Person Name:", t.name, ",Person height:", t.height)
	g := group{"LINE", person{name: "Emily", height: 158}}
	fmt.Println("Group Person Name:", g.person.name, ",Group Height:", g.height, ",Group Name:", g.name)

	t.Greeting()

	g.Greeting()

	//Interface
	u := circle{2}
	k := square{5}
	fmt.Println("Circle area: ", u.area())
	fmt.Println("Square area: ", k.area())

	info(u)
	info(k)

	//Goroutines

	wg.Add(4)
	go foo5()
	go bar2()

	//Race Condition
	fmt.Println("We have $1500")
	go withdraw()
	go withdraw()

	wg.Wait()

	//channel
	// 建立一個有緩衝區大小為3的通道
	ch := make(chan int, 3)

	// 向通道中發送數據
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Sent:", i)
		}
		close(ch)
	}()

	// 從通道中接收數據
	go func() {
		for v := range ch {
			fmt.Println("Received:", v)
			time.Sleep(time.Second)
		}
	}()

	// 等待一段時間以確保所有操作完成
	time.Sleep(6 * time.Second)

	// 使用slice作為列表
	list := []int{0, 1, 2, 3, 4}

	// 打印列表中的元素
	for _, v := range list {
		fmt.Println("List element:", v)
	}
}

// Callback實例
func compile(code []string, callback func(string)) {
	for _, n := range code {
		callback(n)
	}
}

func call1(name string, time string) {
	fmt.Println("Calling " + name + " at " + time + "... ")
}

func call2(name string, time string) (str string) {
	str = "Calling " + name + " at " + time + "... "
	return
}

func plusminas(x, y int) (int, int) {
	return x + y, x - y
}

func total(x ...int) int {
	var t int
	for _, n := range x {
		t += n
	}
	fmt.Println("Print Data...")
	fmt.Println(len(x))
	return t
}

// function回傳function(用法不明)
func foo1() func() string {
	return func() string {
		return "I don't know what is this mean"
	}
}

// 延後執行
func project() {
	defer fmt.Println("Use Try catch")
	fmt.Println("Project Initial")
}

// 傳入記憶體位址
func change(x *string) {
	*x = "Tom"
}

// 自訂Struct
type stuff struct {
	name  string
	price int
}

// 傳入struct
func inprice(s *stuff) {
	s.price += 10
}

func foo3(x int) {
	fmt.Println("[傳值]function記憶體位址", &x)
}

func foo4(y *int) {
	fmt.Println("[傳指標]function記憶體位址", y)
}

type person struct {
	name   string
	height int
}

type group struct {
	name string
	person
}

func (p person) Greeting() {
	fmt.Println("Hi~", p.name)
}

func (g group) Greeting() {
	fmt.Println("We are group", g.name)
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(x shape) {
	fmt.Println("This is a ", x)
	fmt.Println("This area are ", x.area())
}

func foo5() {
	fmt.Println("foo5 Function Start")
	for i := 1; i <= 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}

func bar2() {
	fmt.Println("bar2 Function Start")
	for i := 1; i <= 10; i++ {
		fmt.Println("Bar", i)
	}
	wg.Done()
}

func withdraw() {
	mu.Lock()
	balance := money
	time.Sleep(3000 * time.Millisecond)
	balance -= 1000
	money = balance
	mu.Unlock()
	fmt.Println("After withdrawing $1000, balance: ", money)
	wg.Done()
}
