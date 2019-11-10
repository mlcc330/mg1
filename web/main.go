package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 处理http请求,访问217.0.0.1:8080时,返回"Hello World!!!"
/*
type MyHandler struct {

}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World!!!")
}
*/

// 使用多个处理器,避免了所有请求都返回同一个结果.此时,访问217.0.0.1:8080/hello时,返回"Hello",访问217.0.0.1:8080/world,返回"World"
/*
type HelloHandler struct {

}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello!!!")
}

type WorldHandler struct {

}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "World!!!")
}
*/

// 处理器函数,实现访问217.0.0.1:8080/hello时,返回"Hello",访问217.0.0.1:8080/world,返回"World"
/*
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello!!!")
}

func world(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"World!!!")
}
*/

// 串联两个处理器函数
/*
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello!!!")
}


func log(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - "+name)
		h(w,r)
	}
}
*/

// 请求的header和body
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	//h := r.Header.Get("Accept-Language")
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

// 处理前端templates/index.html页面form表单中的数据
func process(w http.ResponseWriter, r *http.Request) {
	//t, err := templates.ParseFiles("index.html")
	//if err != nil {
	//	log.Fatal("templates.ParseFiles : ", err)
	//	return
	//}
	//t.Execute(w, nil)
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("templates.ParseFiles : ", err)
		return
	}
	t.Execute(w, nil)
}

var dir, sysPath string

func main() {
	// 处理http请求,访问217.0.0.1:8080时,返回"Hello World!!!"
	//handler := MyHandler{}

	// 使用多个处理器,避免了所有请求都返回同一个结果.此时,访问217.0.0.1:8080/hello时,返回"Hello",访问217.0.0.1:8080/world,返回"World"
	//hello := HelloHandler{}
	//world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		// 如果仅使用MyHandler时使用,访问217.0.0.1:8080时,返回"Hello World!!!"
		//Handler: &handler,
	}

	// 使用多个处理器,避免了所有请求都返回同一个结果.此时,访问217.0.0.1:8080/hello时,返回"Hello",访问217.0.0.1:8080/world,返回"World"
	//http.Handle("/hello",&hello)
	//http.Handle("/world",&world)

	// 使用处理器函数,实现访问217.0.0.1:8080/hello时,返回"Hello",访问217.0.0.1:8080/world,返回"World"
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/world", world)

	// 串联两个处理器函数
	//http.HandleFunc("/hello", log(hello))

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/index", requestHandler)
	server.ListenAndServe()
}
