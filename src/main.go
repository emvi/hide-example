package main

import (
	"database/sql"
	"encoding/json"
	"github.com/emvi/hide"
	_ "github.com/lib/pq"
	"log"
)

type Customer struct {
	Id       hide.ID `json:"id"`
	Nullable hide.ID `json:"nullable"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
}

func main() {
	db, _ := sql.Open("postgres", dbString())
	defer db.Close()

	// create a new customer
	customer := Customer{123, 0, "Foobar", 36}

	if _, err := db.Exec(`INSERT INTO "customer" (id, "nullable", "name", age) VALUES ($1, $2, $3, $4)`, customer.Id, customer.Nullable, customer.Name, customer.Age); err != nil {
		panic(err)
	}

	// read customer from database
	rows, err := db.Query(`SELECT * FROM "customer" LIMIT 1`)

	if err != nil {
		panic(err)
	}

	rows.Next()
	if err := rows.Scan(&customer.Id, &customer.Nullable, &customer.Name, &customer.Age); err != nil {
		panic(err)
	}

	// marshal to JSON
	result, err := json.Marshal(&customer)

	if err != nil {
		panic(err)
	}

	log.Println(string(result))

	// unmarshal from JSON
	if err := json.Unmarshal(result, &customer); err != nil {
		panic(err)
	}

	log.Println(customer)

	// parse hash to ID from string
	customerJson := &struct {
		Id string `json:"id"`
	}{}

	if err := json.Unmarshal(result, &customerJson); err != nil {
		panic(err)
	}

	log.Println(hide.FromString(customerJson.Id))
	log.Println(hide.ToString(hide.ID(123)))
}

func dbString() string {
	return "host=localhost" +
		" port=5432" +
		" user=postgres" +
		" password=postgres" +
		" dbname=hide" +
		" sslmode=disable"
}
