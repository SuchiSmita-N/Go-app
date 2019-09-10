package main

import "fmt"

func iterate(character string) string {
    var iterated string
    for i := 0; i < 5; i++ {
    	iterated += character
    }
    return iterated
}

func main() {
        fmt.Println(iterate("S"))
}

