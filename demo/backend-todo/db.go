package main

import (
  "database/sql"
  "log"
  "fmt"
  "os"
  "time"
  "sync"

  _ "github.com/lib/pq"
)

var DB *sql.DB
var dbReady bool
var dbMutex sync.RWMutex

func getEnv(key, fallback string) string {
  val := os.Getenv(key)
  if val == "" {
    return fallback
  }
  return val
}

func InitDB() {
  host := getEnv("DB_HOST", "localhost")
  port := getEnv("DB_PORT", "5432")
  user := getEnv("DB_USER", "postgres")
  password := getEnv("DB_PASSWORD", "postgres")
  dbname := getEnv("DB_NAME", "tododb")
  sslmode := getEnv("DB_SSLMODE", "disable")

  connStr := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode,
  )

  var err error

  maxAttempts := 10
  baseDelay := 1 * time.Second
  maxDelay := 10 * time.Second

  for attempt := 1; attempt <= maxAttempts; attempt++ {
    DB, err = sql.Open("postgres", connStr)

    err = DB.Ping()
    if err == nil {
      dbMutex.Lock()
      dbReady = true
      dbMutex.Unlock()

      log.Println("DB connected successfully")

      // create table if not exists
      createTable()
      return
    }

    log.Printf(
      "[DB INIT] attempt %d/%d failed: %v",
      attempt, maxAttempts, err,
    )

    // exponential backoff: 1s, 2s, 4s, 8s, 10s (cap)
    delay := baseDelay * time.Duration(1<<uint(attempt-1))
    if delay > maxDelay {
      delay = maxDelay
    }

    time.Sleep(delay)
  }


  if err != nil {
    log.Fatal("Database connection issues: ", err)
  }
}

func createTable() {
  query := `
  CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    text TEXT,
    done BOOLEAN DEFAULT FALSE
  )`
  _, err := DB.Exec(query)
  if err != nil {
    log.Fatal(err)
  }
}
