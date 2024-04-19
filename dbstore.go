package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBStore interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	getAccountById(int) (*Account, error)
	updateAccount(*Account) error
}

type PgDBStore struct {
	db *sql.DB
}

type PgClient struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func NewPgClient() (*PgClient, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// conn := "postgres://postgres:postgres@localhost/postgres"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err

	}
	return &PgClient{
		db: db,
	}, nil
	// handling thee config for postgre

}

func (db *PgClient) Initialize() error {
	return db.CreateAccountTable()
}
func (db *PgClient) CreateAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial, 
		created_at timestamp
	)`

	_, err := db.db.Exec(query)
	return err

}

func (db *PgClient) updateAccount(*Account) error {
	return nil
}
func (db *PgClient) getAccountById(id int) error {
	return nil
}

func (db *PgClient) DeleteAccount(id int) error {
	return nil
}
