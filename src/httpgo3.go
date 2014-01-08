/**
 * Name:
 * Comment:
 * Author: Rhinux 
 * Web: http://www.rhinux.info./
 * Created: 2014-01-08 15:54:51
 * Last-Modified: 2014-01-08 15:54:51
 */
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)
var mux map[string]func(http.ResponseWrite, *http.Request)

func main(){
    server := http.Server{
        Addr:     ":8080",
        Handler:   &myHandler{},
        ReadTimeout: 5 * time.Second,
    }
}
mux := make(map[string]func(http.ResponseWrite, *http.Request))
mux["/hello"] = sayHello
mux["/bye"] = sayBye

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "URL"+r.URL.String())
}

