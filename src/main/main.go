package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/m0a/easyjson"
)

func run() error {
    resp, err := http.Get("https://teratail.com/api/v1/questions")
    if err != nil {
        return fmt.Errorf("Failed to connect teratail.com")
    }
    defer resp.Body.Close()

    jsonData, err := easyjson.NewEasyJson(resp.Body)
    if err != nil {
        return fmt.Errorf("Invalid responses")
    }

    fmt.Println(jsonData.K("questions", 0))
    for k, v:=range jsonData.K("questions").RangeObjects() {
        fmt.Printf("key: %d, value: %s\n", k, v.K("title"))
    }

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}

