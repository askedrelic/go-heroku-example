package main

import (
    "os"
    "github.com/hoisie/web"
)

func hello(val string) string { return "hello " + val }

func main() {
    port := os.Getenv("PORT")
    if port == "" { port = "9999" }
    web.Get("/(.*)", hello)
    web.Run(":"+port)
}

