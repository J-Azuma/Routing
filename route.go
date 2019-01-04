package main

import (
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	h := `
	<html><head><title>Hello</title></head><body>
	 Hello
	</body></html>
	`
	fmt.Println(w, h)
}

func goodbyeHandle(w http.ResponseWriter, r *http.Request) {
	h := `
	<html><head><title>Goodbye</title></head><body>
	Goodbye
	</body></html>
	`
	fmt.Println(w, h)
}

func main() {
	//URLごとに関数を登録。
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/goodbye", goodbyeHandle)

	//webサーバーを起動
	if err := http.ListenAndServe(":4000, nil"); err != nil {
		fmt.Println(err)
	}
}
