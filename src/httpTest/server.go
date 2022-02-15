package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		reqBody, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("read req body failed. err=", err.Error())
		}
		fmt.Printf("req body: %s\n", string(reqBody))

		fmt.Fprintln(w, req.URL.String())
		fmt.Fprintln(w, req.Header.Get("Content-Length"))
	})

	err := http.ListenAndServe("127.0.0.1:8001", nil)
	if err != nil {
		fmt.Println("server failed. err=", err.Error())
	}
}
