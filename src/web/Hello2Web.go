package main

import (
    "fmt"
    "net/http"
)

type Hello struct{}

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}
