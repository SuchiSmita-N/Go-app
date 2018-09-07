package main

import "fmt"

const helloPrefix = "Hello, "

func firstgo(name string) string {
    return helloPrefix + name
}
func main() {
        fmt.Println(firstgo("suchi"))
}
