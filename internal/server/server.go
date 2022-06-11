// Package server contains all the logic needed to spin up a new gRPC/HTTP
// server.
package server

import (
	"errors"
)

var (
	// ErrNoDatabase is returned if NewServer is used with a nil database
	// argument.
	ErrNoDatabase = errors.New("no database")
)

// Server is the structure used to create a working server. Obtain a working
// instance with NewServer.
type Server struct {
	db Database
}

// NewDatabase returns a usable instance of the Database structure.
func NewServer(db Database) (srv *Server, err error) {
	if db == nil {
		return nil, ErrNoDatabase
	}
	srv = &Server{
		db: db,
	}

	return srv, nil
}

func (s *Server) unexportedMethod() bool {
	return true
}
