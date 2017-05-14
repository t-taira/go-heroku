package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

// Handler ハンドラー
func Handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}
