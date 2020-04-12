package main

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
)

type Publisher struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var Db *sql.DB

func initDb() {
	var err error
	Db, err = sql.Open("postgres", "host=arjuna.db.elephantsql.com user=yjnwdibd dbname=yjnwdibd password=dcDGXTZeORGq77WsBPkRZ9BvFAYiK8_N sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrievePublisher(id string) (p Publisher, err error) {
	p   = Publisher{}
	err = Db.QueryRow("select id, name, created_at, updated_at from publisher where id = $1", id).Scan(&p.Id, &p.Name, &p.CreatedAt, &p.UpdatedAt)
	return
}

// Create a new publisher
func (publisher *Publisher) create() (err error) {
	statement := "INSERT INTO publisher (name) VALUES ($1) RETURNING id, created_at, updated_at;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(publisher.Name).Scan(&publisher.Id, &publisher.CreatedAt, &publisher.UpdatedAt)
	return
}
