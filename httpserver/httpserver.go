package main

import (
	"fmt"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) () {
	fmt.Println(request.RequestURI)
	writer.Write([]byte("ok"))
}
func main() {
	_ = http.ListenAndServe(":8081", &handler{})
}
