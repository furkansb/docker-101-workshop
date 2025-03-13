package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from a single-stage Dockerfile Go app!\n")
    })

    log.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
