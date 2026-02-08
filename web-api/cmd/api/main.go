// web-api/cmd/api/main.go
package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the Shapes API"))
}

func health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Server is running"))
}

func about(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("John Doe")) // CHANGE TO YOUR NAME
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
    w.Write([]byte(currentTime))
}

func random(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    rand.Seed(time.Now().UnixNano())
    randomNum := rand.Intn(100) + 1
    message := fmt.Sprintf("Random number between 1 and 100: %d", randomNum)
    w.Write([]byte(message))
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/health", health)
    http.HandleFunc("/about", about)
    http.HandleFunc("/time", timeHandler)
    http.HandleFunc("/random", random)

    fmt.Println("Starting Shapes API server on port 4000...")
    fmt.Println("Available routes:")
    fmt.Println("  GET /      - Welcome message")
    fmt.Println("  GET /health - Health check")
    fmt.Println("  GET /about  - About (your name)")
    fmt.Println("  GET /time   - Current server time")
    fmt.Println("  GET /random - Random number")
    
    http.ListenAndServe(":4000", nil)
}