package main

import "fmt"

func main() {
    fmt.Println(say("bah"))
}

func say(msg string) string{
    fmt.Println(msg)
    return msg
}