package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var (
	dbReady bool
	dbMutex sync.RWMutex
)

func getEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}

func Init() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "todo")
	sslmode := getEnv("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)

	maxAttempts := 10
	baseDelay := 1 * time.Second
	maxDelay := 10 * time.Second

	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {

		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf(
				"[DB INIT] sql.Open failed attempt %d/%d: %v",
				attempt,
				maxAttempts,
				err,
			)
			continue
		}

		// connection pool tuning
		DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(25)
		DB.SetConnMaxLifetime(5 * time.Minute)

		err = DB.Ping()
		if err == nil {

			dbMutex.Lock()
			dbReady = true
			dbMutex.Unlock()

			log.Println("DB connected successfully")

			return
		}

		log.Printf(
			"[DB INIT] attempt %d/%d failed: %v",
			attempt,
			maxAttempts,
			err,
		)

		delay := baseDelay * time.Duration(1<<uint(attempt-1))

		if delay > maxDelay {
			delay = maxDelay
		}

		time.Sleep(delay)
	}

	dbMutex.Lock()
	dbReady = false
	dbMutex.Unlock()

	log.Fatal("Database connection failed after retries: ", err)
}

func IsReady() bool {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	return dbReady
}

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("migration complete")
}
