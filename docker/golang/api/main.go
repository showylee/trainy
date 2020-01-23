package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Item struct {
	Id   int
	Name string
}

func main() {
	var Db *sql.DB
	Db, err := sql.Open("postgres", "host=postgres user=gopher password=postgres dbname=go_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	sql := "SELECT id, name FROM ITEM WHERE id = $1;"

	pstatement, err := Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	query := 1
	var item Item

	err = pstatement.QueryRow(query).Scan(&item.Id, &item.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(item.Id, item.Name)
}
