package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"  // 为什么用_ , 在讲解http包时有解释。
)


func myFunc(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hi")
}

func main()  {
	http.HandleFunc("/",myFunc)
	http.ListenAndServe(":8080",nil)
}