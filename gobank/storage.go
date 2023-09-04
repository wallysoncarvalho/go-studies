package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	GetAccountByID(id int) (*Account, error)
	UpdateAccount(account *Account) error
	GetAccounts() ([]*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres password=postgres dbname=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	query := `insert into account values (default, $1, $2, default, 0, now())`

	_, err := s.db.Query(query, account.FirstName, account.LastName)

	return err
}

func (s *PostgresStorage) DeleteAccount(id int) error {

	return nil
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {

	return nil, nil
}

func (s *PostgresStorage) UpdateAccount(account *Account) error {

	return nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	query := "select * from account;"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()

}

func (s *PostgresStorage) createAccountTable() error {
	query := `create table  if not exists account(
		id serial primary key,
		first_name varchar not null,
		last_name varchar not null,
		number serial,
		balance numeric(1000,2) not null constraint positive_balance check(balance >= 0),
		created_at timestamp with time zone
	)`

	_, err := s.db.Exec(query)

	return err
}
