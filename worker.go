package main

import (
    "fmt"
)

func main() {

    messages := make(chan string)

    go work(messages)

    msg := <-messages
    fmt.Println("Channel running on", msg)
}

func work(messages chan<- string) {
    messages <- "golangcode.com"
}