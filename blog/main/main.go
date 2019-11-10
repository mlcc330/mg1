package main

import (
	"fmt"
	"mg1/blog/dbfunc"
)

func main() {
	fmt.Println("----------------我的博客练习应用---------------------")
	dbfunc.ConnMongoDB()
}
