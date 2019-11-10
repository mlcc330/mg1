package main

import "fmt"

/*
* Author : mlcc
* Create Time :  上午11:47
* Description : 数组
* 在典型的 Go 程序中，相对于数组而言，slice 使用的更多
 */

func main() {
	var arr [5]int //创建了一个数组 arr 来存放刚好 5 个 int
	fmt.Println("emp arrays : ", arr)

	arr[4] = 100
	fmt.Println("set:", arr)      // 使用array[index] = value 语法来设置数组指定位置的值
	fmt.Println("get:", arr[4])   // 用array[index] 得到值
	fmt.Println("len:", len(arr)) // 使用内置函数 len 返回数组的长度

	b := [5]int{1, 2, 3, 4, 5} // 使用这个语法在一行内初始化一个数组
	fmt.Println("dcl:", b)

	fmt.Println("组合数据来构造多维的数据结构")
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
