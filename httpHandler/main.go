package main

import "net/http"

type fooHandler struct {
	Message string
}

// 对于Handle，我们需要定义一个handler类型，然后需要这个类型去实现ServeHTTP
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

//对于HandleFunc， 只需要传入一个func符合func(ResponseWriter, *Request)即可
func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

func main() {

	http.HandleFunc("/bar", barHandler)
	http.Handle("/foo", &fooHandler{Message: "foo called"})
	http.ListenAndServe(":5000", nil)
}
