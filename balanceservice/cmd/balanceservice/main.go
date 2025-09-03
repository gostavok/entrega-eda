package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
	"fmt"
	"strings"
	"encoding/json"
	"github.com/gostavok/entrega-eda/balanceservice/internal/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting Balance Service on :3003...")

	// Database connection
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/balances", dbUser, dbPass, dbHost, dbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Start Kafka consumer
	kafka.StartBalanceConsumer(db)

	http.HandleFunc("/balances/", func(w http.ResponseWriter, r *http.Request) {
		accountID := strings.TrimPrefix(r.URL.Path, "/balances/")
		if accountID == "" {
			http.Error(w, "account_id required", http.StatusBadRequest)
			return
		}
		var balance float64
	err := db.QueryRow("SELECT balance FROM balances WHERE account_id = ?", accountID).Scan(&balance)
		if err == sql.ErrNoRows {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		resp := map[string]interface{}{"account_id": accountID, "balance": balance}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	log.Fatal(http.ListenAndServe(":3003", nil))
}
