package main

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前 arr = \n", (*arr))
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后 arr = \n", (*arr))
}

func Sort(arr *[5]int) {
	fmt.Println("排序前 arr = \n", (*arr))
	temp := 0 // 临时变量(用于做交换)

	// 完成第1轮排序(外层1次)
	for j := 0; j < 4; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换
			temp := (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第1次排序后 arr = \n", (*arr))

	// 完成第2轮排序(外层2次)
	for j := 0; j < 3; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换
			temp := (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第2次排序后 arr = \n", (*arr))

	// 完成第3轮排序(外层3次)
	for j := 0; j < 2; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换
			temp := (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第3次排序后 arr = \n", (*arr))

	// 完成第4轮排序(外层4次)
	for j := 0; j < 1; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换
			temp := (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第4次排序后 arr = \n", (*arr))
}

func main() {
	// 定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	// 将数组传递给一个函数,完成排序
	BubbleSort(&arr)
}
