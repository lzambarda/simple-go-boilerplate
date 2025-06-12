// Package database contains all the logic needed to interact with the database.
//
// Why isn't this package called "db"? That's because we want to use "db" as the
// shorthand of an instance of database, and this saves some head-scratching.
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Database is the structure used to interact with the database. Obtain a working
// instance with NewDatabase.
//
// This has to be seen as an implementation of any other place in the repo which
// needs a database. The other packages will work with their internal
// "contracts" and this struct must implement them. Check out server/contract.go
// for more details.
type Database struct {
	// An example db handle, can use anything, even github.com/jackc/pgx
	handle sql.DB
}

// New returns a usable instance of the Database structure.
func New() (*Database, error) {
	return nil, errors.New("not implemented") //nolint:err113 // Fine here.
}

// DoSomething does what is written on the tin.
func (db *Database) DoSomething() (string, error) {
	var res string

	err := db.handle.QueryRow(queries[queryFooName]).Scan(&res)
	if err != nil {
		return "", fmt.Errorf("query row: %w", err)
	}

	return res, nil
}
