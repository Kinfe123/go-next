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
func (db *PgClient) GetAccountById(id int) (*Account, error) {
	query := `select * from account where id = $1`
	account := Account{}
	row := db.db.QueryRow(query, id)
	err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.AccountNumber, &account.Balance, &account.Created_at)

	if err != nil {
		return nil, err
	}
	return &account, nil

}

func (db *PgClient) DeleteAccount(id int) error {
	query := `delete from account where id = $1`
	_, err := db.db.Exec(query, id)
	return err

}

func (db *PgClient) GetAllAccount() error {
	query := `select * from account`
	_, err := db.db.Query(query)
	return err
}

func (db *PgClient) CreateAccount(acc *Account) error {
	query := `insert into account (first_name , last_name , number  ,  balance , created_at) values ($1 , $2 , $3 , $4 , $5)`
	_, err := db.db.Exec(query, acc.FirstName, acc.LastName, acc.AccountNumber, acc.Balance, acc.Created_at)
	return err

}

func (db *PgClient) deleteAllAccount() error {
	query := `delete from account;`
	_, err := db.db.Exec(query)
	return err

}

func (db *PgClient) SelectAllAccount() ([]*Account, error) {
	rows, err := db.db.Query(`select * from account`)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := Account{}
		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.AccountNumber, &account.Balance, &account.Created_at); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (db *PgClient) CheckSenderBalance(fromId int) (int64, error) {
	account, err := db.GetAccountById(fromId)
	if err != nil {
		return -1, nil
	}
	return account.Balance, nil

}

func (db *PgClient) Debit(fromId int, amount int) error {
	query := `update account set balance = balance - $1 where id = $2`
	_, err := db.db.Exec(query, amount, fromId)
	return err

}
func (db *PgClient) Credit(toId int, amount int) error {


	query := `update account set balance = balance + $1 where id = $2`
	_, err := db.db.Exec(query,amount , toId)
	return err

}

func (db *PgClient) Withdraw(fromId int , amount int) error {
    err := db.Debit(fromId , amount)
	return err
}