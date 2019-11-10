package charpt3

import "fmt"

var arr1 = [5]int{10, 20, 30, 40, 50}
var arr2 = [...]int{10, 20, 30, 40, 50}
var arr3 = [5]int{1: 10, 2: 20}

var slice1 = []int{10, 20, 30, 40, 50}
var slice2 = []string{"Red", "Blue", "Green", "Yellow", "Pink"}
var slice3 = make([]int, 3)

func PrintArrys() {
	fmt.Println("arr1 is : ", arr1)
	fmt.Println("arr2 is : ", arr2)
	fmt.Println("arr3 is : ", arr3)
}

func PrintSlice() {
	fmt.Println("slice1 : ", slice1)
	fmt.Println("slice2 : ", slice2)
	fmt.Println("slice2[2,3,5] is : ", slice2[0:3:3])
	fmt.Println("slice3 : ", slice3)
	fmt.Println("append(slice1, 60) is : ", append(slice1, 60))

	fmt.Println("-----------range迭代切片中的数据--------------")
	for index, value := range slice1 {
		fmt.Printf("Index : %d, Value : %d \n", index, value)
	}

	fmt.Println("-----------range迭代切片中的数据,忽略那些不重要的数据---------------------")
	for _, value := range slice1 {
		fmt.Printf("Value is : %d \n", value)
	}

	fmt.Println("-----------range迭代切片中的数据,但都是从0开始,此刻,用for来实现从特定的开始循环-------------------")
	for index := 2; index < len(slice1); index++ {
		fmt.Printf("Index is %d,  Value is : %d \n", index, slice1[index])
	}
}
