package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Item struct {
	Id   int
	Name string
}

func main() {
	e := echo.New()
	routing(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func routing(e *echo.Echo) {
	e.GET("/", healthcheck)
	e.GET("/test", test)
}

func test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "test"})
}

func healthcheck(c echo.Context) error {
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
	return c.JSON(http.StatusOK, map[string]string{"id": strconv.Itoa(item.Id), "name": item.Name})
}
