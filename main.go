package main

import (
  "net/http"
  "fmt"
)

func main() {
  db := createDatabase()
  server := newServer(db)

  fmt.Println("Starting server on :8080")
  http.ListenAndServe(":8080", server)
}
