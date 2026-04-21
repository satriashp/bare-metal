package main

import (
  "log"
  "net/http"
)

func main() {
  go InitDB() //run in background

  http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      GetTodos(w, r)
    case http.MethodPost:
      CreateTodo(w, r)
    default:
      http.Error(w, "method not allowed", 405)
    }
  })

  http.HandleFunc("/healthz", Healthz)
  http.HandleFunc("/readyz", Readyz)

  http.HandleFunc("/debug", Debug)

  log.Println("Server running on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}