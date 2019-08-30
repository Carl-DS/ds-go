package main

import "fmt"

func main() {
	// map 是引用类型，可如下声明：
	// 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。
	// 未初始化的 map 的值是 nil

	// map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。通过
	// key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果
	// 你很在乎性能的话还是建议用切片来解决问题

	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415926
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is:%d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is:%f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is:%d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is:%d\n", mapLit["ten"])

	// map 是 引用类型 的： 内存用 make 方法来分配。
	// map 的初始化： var map1 = make(map[keytype]valuetype) 。
	// 或者简写为： map1 := make(map[keytype]valuetype)
	// 不要使用 new，永远用 make 来构造 map
	// 注意 如果你错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并
	// 且取了它的地址

	// 为了说明值可以是任意类型的，这里给出了一个使用 func() int 作为值的 map：
	// 输出结果为： map[1:0x10903be0 5:0x10903ba0 2:0x10903bc0] : 整形都被映射到函数地址
	mf := map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		},
	}
	fmt.Println(mf)

	// 和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是你也可以选择
	//标明 map 的初始容量 capacity ，就像这样： make(map[keytype]valuetype, cap) 。例如：
	map2 := make(map[string]float32, 100)
	map2["keyMap2"] = 3.25
	// 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大
	// 的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明
	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440,
	}
	fmt.Println(noteFrequency)

	sliceMap1 := make(map[int][]int)
	sliceMap2 := make(map[int]*[]int)
	sliceMap1[1] = []int{1, 2, 3}
	fmt.Println(sliceMap1)
	sliceMap2[1] = &[]int{4, 5, 6}
	fmt.Println(sliceMap2)

	val2, isPresent := map2["keyMap22"]
	if !isPresent {
		fmt.Printf("map2 中不存在key：keyMap22\n")
	}
	fmt.Println(val2)
}
