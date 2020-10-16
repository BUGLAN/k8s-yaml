package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)
	fmt.Println("server start")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "pong")
}
