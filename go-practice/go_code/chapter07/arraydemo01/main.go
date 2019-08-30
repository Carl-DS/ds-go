package main

import (
	"fmt"
)

func main() {
	hen1 := 3.0
	hen2 := 4.0
	hen3 := 5.0
	hen4 := 6.6

	totalWeight := hen1 + hen2 + hen3 + hen4
	avgWeight := totalWeight / 4
	fmt.Printf("totalWeight = %v, avgWeight = %.2f \n", totalWeight, avgWeight)
	avgWeight2 := fmt.Sprintf("%.2f", totalWeight/4)
	fmt.Printf("totalWeight = %v, avgWeight = %v \n", totalWeight, avgWeight2)

	// 1. 定义一个数组
	var hens [4]float64
	// 2. 给数组的每个元素赋值, 元素的下标是从0开始的
	hens[0] = 2.0
	hens[1] = 3.0
	hens[2] = 4.0
	hens[3] = 5.0
	// 3. 遍历数组求出总体量
	var total = 0.0
	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}
	// 4. 求出平均体重
	avg := fmt.Sprintf("%.2f", total/float64(len(hens)))
	fmt.Printf("totalWeight = %v, avgWeight = %v \n", total, avg)
}
