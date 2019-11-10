package main

import "fmt"

/*
* Filename : 切片
* Author : mlcc
* Create Time : 2019-05-25 19-5-25
* Description :
 */

func main() {
	//	slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。要创建一个长度非零的空slice，需要使用内建的方法 make。
	s := make([]string, 3) // 创建了一个长度为3的 string 类型 slice（初始化为零值）
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // len 返回 slice 的长度

	// append返回一个包含了一个或者多个新值的 slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 copy。这里我们创建一个空的和 s 有相同长度的 slice c，并且将 s 复制给 c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice 支持通过 slice[low:high] 语法进行“切片”操作
	l := s[2:5] // 得到一个包含元素 s[2], s[3],s[4] 的 slice
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l) // 从 s[0] 到（但是包含）s[5]
	// 这个 slice 从（包含）s[2] 到 slice 的后一个值。

	l = s[2:] // 从（包含）s[2] 到 slice 的后一个值
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // 在一行代码中声明并初始化一个 slice 变量
	fmt.Println("dcl:", t)

	fmt.Println("可以组成多维数据结构。内部的 slice 长度可以不同")
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		//fmt.Println("innerLen",innerLen)
		twoD[i] = make([]int, innerLen)
		//fmt.Println("twoD[i]",twoD[i])
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			//fmt.Println("twoD[i][j]",twoD[i][j])
		}
	}
	fmt.Println("2d: ", twoD)
}
