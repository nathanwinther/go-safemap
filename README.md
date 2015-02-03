# Go programming language concurrent safe map

Concurrent safe map using sync.RWMutex

## Installation

This package can be installed with the go get command:

    go get github.com/nathanwinther/go-safemap

## Usage

    package main

    import (
        "fmt"
        "os"
        "github.com/nathanwinther/go-safemap"
    )

    func main() {
        config := safemap.New()

        config.Set("one", "Hello")
        config.Set("two", "World")

        config.Dump(os.Stdout)

        k := "one"
        fmt.Printf("%s = %s\n", k, config.Get(k))

        k = "two"
        v, ok := config.Test(k)
        if ok {
            fmt.Printf("%s = %s\n", k, v)
        }
    }

