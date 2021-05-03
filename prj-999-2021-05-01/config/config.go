package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	db     *sql.DB
	logger *log.Logger
}

func NewLocal() (*Config, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/goofus")
	if err != nil {
		return nil, err
	}

	c := &Config{
		db:     db,
		logger: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile),
	}

	return c, nil
}

func (c *Config) DB() *sql.DB {
	return c.db
}

func (c *Config) Logger() *log.Logger {
	return c.logger
}
