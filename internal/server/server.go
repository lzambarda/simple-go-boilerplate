// Package server contains all the logic needed to spin up a new gRPC/HTTP
// server.
package server

import (
	"errors"
)

// ErrNoDatabase is returned if NewServer is used with a nil database
// argument.
var ErrNoDatabase = errors.New("no database")

// Server is the structure used to create a working server. Obtain a working
// instance with NewServer.
type Server struct {
	db Database
}

// New returns a usable instance of the Database structure.
func New(db Database) (*Server, error) {
	if db == nil {
		return nil, ErrNoDatabase
	}

	srv := &Server{
		db: db,
	}

	return srv, nil
}

//nolint:unused // Fine here.
func (s *Server) unexportedMethod() bool {
	return true
}
