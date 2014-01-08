/**
 * Name:
 * Comment:
 * Author: Rhinux 
 * Web: http://www.rhinux.info./
 * Created: 2014-01-08 15:09:40
 * Last-Modified: 2014-01-08 15:10:07
 */
package main

import (
    "io"
    "log"
    "net/http"
)

func main(){
    http.HandleFunc("/", sayHello)

    err := http.ListenAndServe(":8080",nil)
    if err != nil{
        log.Fatal(err)
    }
}

func sayHello(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello world, this is version 1.")
}