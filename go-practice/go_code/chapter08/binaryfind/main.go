package main

import (
	"fmt"
)

// 二分查找的函数
/*
二分查找的思路: 比如我们要查找的是 findVal
1. arr 是一个有序数组,并且是从小到大排序
2. 先找到中间的下标 middle = (leftIndex + rightIndex) / 2, 然后让中间下标的值和findVal进行比较
2.1 如果 arr[middle] > findVal, 就应该向 leftIndex --- (middle - 1) 方向找
2.2 如果 arr[middle] < findVal, 就应该向 middle+1 --- rightIndex
2.3 如果 arr[middle] = findVal, 就找到
2.4 上面的2.1, 2.2, 2.3 的逻辑会递归执行
3. 什么情况下, 说明找不到
if leftIndex > rightIndex {
	// 找不到
	return
}
*/
func BinaryFind(arr *[6]int, leftIndex, rightIndex, findVal int) {
	// 判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		// 说明我们要查找的数应该在 leftIndex --- (middle - 1) 方向找
		BinaryFind(arr, leftIndex, (middle - 1), findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, (middle + 1), rightIndex, findVal)
	} else {
		fmt.Printf("找到了, 下标为: %v \n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 12324)
}
