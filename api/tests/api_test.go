package tests

import (
	"encoding/json"
	"fmt"
	c "generate/bootcamp/src/controller"
	"generate/bootcamp/src/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/huandu/go-assert"
	"github.com/jackc/pgx"
)

func TestGetBooks(t *testing.T) {
	db_url, exists := os.LookupEnv("DATABASE_URL")

	cfg := pgx.ConnConfig{
		User:     "postgres",
		Database: "backendbootcamp",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     5432,
	}
	var err error
	if exists {
		cfg, err = pgx.ParseConnectionString(db_url)

		if err != nil {
			panic(err)
		}
	}
	conn, err := pgx.Connect(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	m := &model.PgModel{
		Conn: conn,
	}
	c := &c.PgController{
		Model: m,
	}
	router := c.Serve()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/books/1738", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books model.Book

	if e := json.Unmarshal(w.Body.Bytes(), &books); e != nil {
		panic(err)
	}

	test_book := model.Book{
		BookId: 1738,
		Title:  "The Lightning Thief",
		Author: "Rick Riordan",
	}
	assert.Equal(t, test_book, books)
}
