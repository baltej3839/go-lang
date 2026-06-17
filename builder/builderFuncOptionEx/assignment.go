package main

import "fmt"

type Database struct {
	Host           string
	Port           int
	SSL            bool
	MaxConnections int
}

type Option func(*Database)

func NewDatabase(opts ...Option) *Database {
	db := &Database{
		Host:           "localhost",
		Port:           5432,
		SSL:            false,
		MaxConnections: 10,
	}

	for _, opt := range opts {
		opt(db)
	}

	return db
}

func WithHost(newHost string) Option {
	return func(db *Database) {
		db.Host = newHost
	}
}

func main() {

	db := NewDatabase(WithHost("new user"))
	fmt.Println(db.Host)
}
