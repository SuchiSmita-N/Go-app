package main

import "fmt"

const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const hindiHelloPrefix = "नमस्ते, "
const hindi = "Hindi"
const spanish = "Spanish"

func firstgo(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case hindi:
        prefix = hindiHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishPrefix
    }
    return
}
func main() {
        fmt.Println(firstgo("suchi", "Hindi"))
}

