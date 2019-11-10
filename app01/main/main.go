package main

import (
	"html/template"
	"log"
	_ "mg1/app01/utils"
	"net/http"
	"os"
	//"nlu/log"
)

func init() {
	log.SetOutput(os.Stdout)
}

type Todo struct {
	Task string
	Done bool
}

func main() {
	//fmt.Print("Hello World!!!")
	//search.Run("President")

	//fmt.Println("---------------charpt3 package------------------")
	//charpt3.PrintArrys()
	//charpt3.PrintSlice()

	//fmt.Println("---------------解析toml配置文件------------------")
	//flag.Parse()
	//if err := utils.Init(); err != nil{
	//  //log.Error("conf.Init() err:%+v", err)
	//  fmt.Println("err is : ",err)
	//}
	//
	//mysqlconf := utils.Conf.Mysql
	//fmt.Println(mysqlconf.Dbname)

	tmpl := template.Must(template.ParseFiles("/home/mlcc/dev/GOPATH/src/mg1/app01/html/todos.html"))
	todos := []Todo{
		{"Learn Go", true},
		{"Read Go Web Example", true},
		{"Create a web app in Go", true},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl.Execute(writer, struct{ Todo []Todo }{todos})
	})

	http.ListenAndServe(":8080", nil)
}
