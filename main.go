package main

import "fmt"

func main() {
    fmt.Println("[fractalscope] initializing system")

    for i := 0; i < 12; i++ {
        fmt.Println("[cycle]", i)
    }

    fmt.Println("[fractalscope] shutdown complete")
}
