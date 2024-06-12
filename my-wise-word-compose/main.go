package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    router := gin.Default()

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
        dbUser, dbPass, dbHost, dbPort, dbName)

    fmt.Println("Connect to: ", connStr)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }

    defer db.Close()

    router.GET("/", func(c *gin.Context) {
		c.String(200, "Ouch!")
		fmt.Println("To Infinity and Beyond!")
	})

    router.GET("/ping", func(c *gin.Context) {
        if err := db.Ping(); err != nil {
            fmt.Println("Error pinging database:", err)
            c.String(500, "Database connection failed")
        } else {
            fmt.Println("Ping to database successful")
            c.String(200, "Ping to Database connection successful")

            // Insert current timestamp
            timestamp := time.Date(2024, 1, 31, 19, 55, 0, 0, time.UTC)
            _, err := db.Exec("INSERT INTO access_log (timestamp) VALUES ($1)", timestamp)
            if err != nil {
                fmt.Println("Error inserting timestamp:", err)
            } else {
                fmt.Println("Inserted timestamp:", timestamp)
            }
        }
    })

    err = router.Run(":78")
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
