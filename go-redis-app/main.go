package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/redis/go-redis/v9"
    "context"
)

var ctx = context.Background()

func main() {
    // Environment variables for port & Redis address
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr == "" {
        redisAddr = "redis:6379" // Docker Compose service name
    }

    rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
    log.Printf("Connected to Redis at %s\n", redisAddr)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        count, err := rdb.Incr(ctx, "counter").Result()
        if err != nil {
            http.Error(w, "Error incrementing counter", http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Hello from Dockerized Go + Redis!\n")
        fmt.Fprintf(w, "This page has been viewed %d times.\n", count)
    })

    log.Printf("Server running on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
