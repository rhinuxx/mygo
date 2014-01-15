/**
 * Name:
 * Comment:
 * Author: Rhinux
 * Web: http://www.rhinux.info./
 * Created: 2014-01-06 13:28:43
 * Last-Modified: 2014-01-09 19:17:40
 */
package main

import (
    "fmt"
    "os"
)

func main() {
    c1, c2 := make(chan int), make(chan string)
    o := make(chan bool, 2)
    go func() {
        for {
            select {
            case v, ok := <-c1:
                if !ok {
                    o <- true
                    break
                }
                fmt.Println("c1", v)
            case v, ok := <-c2:
                if !ok {
                    o <- true
                    break
                }
                fmt.Println("c2", v)
            }
        }
    }()
    c1 <- 1
    c2 <- "hi"
    c1 <- 3
    c2 <- "hello"
    close(c1)
    close(c2)
    <-o
}
