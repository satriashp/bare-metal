package handlers

import (
  "context"
  "sync"
  "database/sql"
  "encoding/json"
  "net/http"
  "time"
  "runtime"
  "fmt"

  "todo-core/internal/repository"
)

var (
	dbReady bool
	dbMutex sync.RWMutex
)

var DB *sql.DB

type CreateTodoRequest struct {
  Text string `json:"text"`
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
  todos, err := repository.GetTodos()
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(todos)
}


func CreateTodo(w http.ResponseWriter, r *http.Request) {
  var req CreateTodoRequest
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    http.Error(w, "invalid request body", 400)
    return
  }

  err = repository.CreateTodo(req.Text)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  w.WriteHeader(http.StatusCreated)
}

// func CreateTodo(w http.ResponseWriter, r *http.Request) {
//   var todo Todo
//   json.NewDecoder(r.Body).Decode(&todo)

//   err := DB.QueryRow(
//     "INSERT INTO todos(text, done) VALUES($1, $2) RETURNING id",
//     todo.Text, todo.Done,
//   ).Scan(&todo.ID)

//   if err != nil {
//     http.Error(w, err.Error(), 500)
//     return
//   }

//   json.NewEncoder(w).Encode(todo)
// }

// func GetTodos(w http.ResponseWriter, r *http.Request) {
//   // --- ADD THIS LOGGING BLOCK ---
//   fmt.Println("--- Incoming Request Headers ---")
//   for name, values := range r.Header {
//       for _, value := range values {
//           fmt.Printf("%s: %s\n", name, value)
//       }
//   }
//   fmt.Println("-------------------------------")
//   // --- END LOGGING BLOCK ---

//   rows, err := DB.Query("SELECT id, text, done FROM todos")
//   if err != nil {
//     http.Error(w, err.Error(), 500)
//     return
//   }
//   defer rows.Close()

//   var todos []Todo
//   for rows.Next() {
//     var t Todo
//     rows.Scan(&t.ID, &t.Text, &t.Done)
//     todos = append(todos, t)
//   }

//   json.NewEncoder(w).Encode(todos)
// }

func Healthz(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("OK: healthy"))
}

func Readyz(w http.ResponseWriter, r *http.Request) {
  dbMutex.RLock()
  ready := dbReady
  dbMutex.RUnlock()

  if !ready {
    http.Error(w, "not ready", http.StatusServiceUnavailable)
    return
  }

  // check db connection
  ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
  defer cancel()

  err := DB.PingContext(ctx)
  if err != nil {
    http.Error(w, "DB connection fail", http.StatusServiceUnavailable)
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte("ready"))
}

func Debug(w http.ResponseWriter, r *http.Request) {
  count := runtime.NumGoroutine()
  w.Write([]byte(fmt.Sprintf("goroutines: %d", count)))
}
