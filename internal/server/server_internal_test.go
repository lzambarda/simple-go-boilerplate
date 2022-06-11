//go:build !integration

// This file can be used to test internal, unexported stuff belonging to the
// server package.
//
// Internal tests can also be used to export something that should only be
// exported during tests. See "GetDB" for an example.
package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServerInternal is a simple internal unit test
func TestServerInternal(t *testing.T) {
	s := Server{}
	expected := true
	actual := s.unexportedMethod()
	assert.Equal(t, expected, actual)
}

// GetDB allows to access the unexported db property, only available while
// running "go test".
func (s *Server) GetDB() Database {
	return s.db
}
