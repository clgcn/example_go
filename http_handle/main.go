package main

import "net/http"

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func fooHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("call foo"))
}
func main() {
	// http.Handle("/foo", &fooHandler{Message: "hello world"})
	http.HandleFunc("/foo", fooHTTP)
	http.ListenAndServe(":5000", nil)
}
