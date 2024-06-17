package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "About Page")
}
