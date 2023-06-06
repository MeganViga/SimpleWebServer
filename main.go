package main

import (
	"fmt"
	"net/http"
)
func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/home"{
		fmt.Fprintf(w,"Not found")
		return
	}
	if r.Method != "GET"{
		fmt.Fprintf(w,"Not Correct Method")
		return
	}
	fmt.Fprintf(w, "hello")
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"err: %v", err)
		return
	}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	fmt.Println(fname, lname)
	fmt.Fprintf(w, fname + lname)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/form",formHandler)
	http.ListenAndServe(":8080", nil)
}