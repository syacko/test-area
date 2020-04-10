package main

import (
    "fmt"
    "rsc.io/quote"
    "github.com/syacko/test-area/slogger"
)

func main() {
    fmt.Println(quote.Hello())
    slogger.Info("Test message")
}
