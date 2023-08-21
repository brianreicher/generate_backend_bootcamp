package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	c "generate/bootcamp/src/controller"
	"generate/bootcamp/src/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/huandu/go-assert"
)

func TestGetBooks(t *testing.T) {
	conn, err := sql.Open("mysql", "remo:pwd@tcp(localhost:3333)/remodb")
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

	req, _ := http.NewRequest("GET", "/v1/books/9780786838653", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books model.Book

	if e := json.Unmarshal(w.Body.Bytes(), &books); e != nil {
		panic(err)
	}

	test_book := model.Book{
		BookId:      "3757",
		Title:       "Percy Jackson and the Olympians",
		Author:      "Rick Riordan",
		ISBN_13:     "9780786838653",
		ISBN_10:     "0786838655",
		Subtitle:    "",
		PublishDate: "2006",
		PageCount:   "375",
		Synopsis:    "Percy Jackson is a good kid, but he can't seem to focus on his schoolwork or control his temper. And lately, being away at boarding school is only getting worse-Percy could have sworn his pre-algebra teacher turned into a monster and tried to kill him. When Percy's mom finds out, she knows it's time that he knew the truth about where he came from, and that he go to the one place he'll be safe. She sends Percy to Camp Half Blood, a summer camp for demigods (on Long Island), where he learns that the father he never knew is Poseidon, God of the Sea. Soon a mystery unfolds and together with his friends -- one a satyr and the other the demigod daughter of Athena -- Percy sets out on a quest across the United States to reach the gates of the Underworld (located in a recording studio in Hollywood) and prevent a catastrophic war between the gods.",
	}
	assert.Equal(t, test_book, books)
}
