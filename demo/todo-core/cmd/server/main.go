package main

import (
  "log"
  "net/http"

  "todo-core/internal/db"
  "todo-core/internal/handlers"
)

func main() {
  db.Init()
  db.Migrate()

  http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      handlers.GetTodos(w, r)
    case http.MethodPost:
      handlers.CreateTodo(w, r)
    default:
      http.Error(w, "method not allowed", 405)
    }
  })

  http.HandleFunc("/healthz", handlers.Healthz)
  http.HandleFunc("/readyz", handlers.Readyz)

  http.HandleFunc("/debug", handlers.Debug)

  log.Println("Server running on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}