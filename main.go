package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/serie-hugokaku/matching/ent"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env: %+v", err)
	}

	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/%s?parseTime=True", pass, dbname)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %+v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Debug().Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema: %+v", err)
	}
}

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server is not running: %+v", err)
	}
}
