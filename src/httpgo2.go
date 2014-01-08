/**
 * Name:
 * Comment:
 * Author: Rhinux 
 * Web: http://www.rhinux.info./
 * Created: 2014-01-08 15:27:58
 * Last-Modified: 2014-01-08 16:05:51
 */
package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", &myHandler{})
    mux.HandleFunc("/hello", sayHello)
    wd, err := os.Getwd()
    if err != nil{
        log.Fatal(err)
    }
    mux.Handle("/static/",http.StripPrefix("/static/",
        http.FileServer(http.Dir(wd))))
    err = http.ListenAndServe(":8080",mux)
    if err != nil {
        log.Fatal(err)
    }
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "URL"+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "hello world version2")
}
