package main

import (
    "os"
    "fmt"
    "net/http"
)

func run() error {
    resp, err := http.Get("https://teratail.com/api/v1/questions")
    if err != nil {
        return fmt.Errorf("Failed to connect teratail.com")
    }
    defer resp.Body.Close()

    fmt.Println(resp)
    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}
