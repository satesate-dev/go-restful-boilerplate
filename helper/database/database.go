package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //pq is a pure Go Postgres driver for the database/sql package
)

//NewDatabase() is used to create new Databse setup
func NewDatabase(db, username, password, host, port, dbName, timezone, sslMode, sslCert, sslKey, sslRootCert string) *Database {
	return &Database{
		db:          db,
		username:    username,
		password:    password,
		host:        host,
		port:        port,
		dbName:      dbName,
		timezone:    timezone,
		sslMode:     sslMode,
		sslCert:     sslCert,
		sslKey:      sslKey,
		sslRootCert: sslRootCert,
	}
}

// Setup struct
type Database struct {
	db          string
	host        string
	dbName      string
	username    string
	password    string
	port        string
	timezone    string
	sslMode     string
	sslCert     string
	sslKey      string
	sslRootCert string
}

//Connect() is used to connect to DB
func (c *Database) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?TimeZone=%s&sslmode=%s", c.db, c.username, c.password, c.host, c.port, c.dbName, c.timezone, c.sslMode,
	)

	if c.sslMode == "require" {
		connStr += fmt.Sprintf("&sslcert=%s&sslkey=%s&sslrootcert=%s", c.sslCert, c.sslKey, c.sslRootCert)
	}

	db, err := sql.Open(c.db, connStr)

	return db, err
}
