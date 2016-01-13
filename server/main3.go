package main

import (
	"io"
	"net/http"
  "fmt"
)

//func hello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, fmt.Sprintf("Hello world! %v", f()))
//}

func closure () func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func createHandler(c func() int) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, fmt.Sprintf("Hello world! %v", c()))
  }
}

func main() {
  c := closure();
	http.HandleFunc("/", createHandler(c))
	http.ListenAndServe(":8003", nil)
}
