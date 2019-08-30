package gointerview

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 写出下面代码输出内容
func Test1(t *testing.T) {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

	// defer 函数属延迟执行,延迟到调用者函数执行return 命令前被执行.多个defer 之前按FILO 后进先出顺序执行
	// 故考题中，在Panic触发时结束函数运行，在return前先依次打印:打印后、打印中、打印前 。最后由runtime运行时抛出打印panic异常信息。
	// 需要注意的是, 函数的return value 不是原子操作,而是在编译器中分解为两部分, 返回值赋值和return .
	// 而defer 刚好被插入到末尾的return 前执行.故可以在derfer 函数中修改返回值

}

// 以下代码有什么问题，说明原因
func Test2(t *testing.T) {

	type student struct {
		Name string
		Age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	fmt.Printf("*student=> %v \n", m)
	// for 遍历时, 变量stu 指针不变, 每次遍历仅进行struct 值拷贝, 故m[stu.Name] = &stu 实际上一致指向同一个指针,
	// 最终该指针的值为遍历的最后一个struct 的值拷贝
	// 修正:取数据组原始值的指针
	for i, _ := range stus {
		stu := stus[i]
		m[stu.Name] = &stu
	}
	fmt.Printf("*student=> %v \n", m)
}

// 下面的代码会输出什么，并说明原因
func Test3(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("one => i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("two => i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 下面代码会输出什么？
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func test4() {
	t := Teacher{}
	t.ShowA()
}

func Test4(t *testing.T) {
	test4()
}

// 5、下面代码会触发异常吗？请详细说明
func Test5(t *testing.T) {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

// 答: 有可能发生异常, 是随机事件

// 6、下面代码输出什么？
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func Test6(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// 7、请写出以下输入内容
func Test7(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 8、下面的代码有什么问题?
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func Test8(t *testing.T) {
	userAges := UserAges{}
	userAges.Add("carl", 25)
}
