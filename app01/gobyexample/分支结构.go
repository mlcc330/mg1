package main

/*
* Filename : ${NAME}
* Author : ${USER}
* Create Time : ${DATA} ${TIME}
* Description :在一个 case 语句中，你可以使用逗号来分隔多个表达式。在这个例子中，我们很好的使用了可选的default 分支。
 */

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Println("write ", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 4:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon.")
	}
}
