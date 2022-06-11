//go:build !integration

// Package server_test contains simple unit and integration tests for exported
// stuff in server.
package server_test

import (
	"testing"

	"github.com/ORGNAME/REPONAME/internal/database"
	"github.com/ORGNAME/REPONAME/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("NewServer", testNewServer)
}

// Simple test case structure
func testNewServer(t *testing.T) {
	tcs := map[string]struct {
		inputDB     server.Database
		expectedErr error
	}{
		"no db": {
			nil,
			server.ErrNoDatabase,
		},
		"ok": {
			&database.Database{},
			nil,
		},
	}
	for name, tc := range tcs {
		tc := tc // avoid referencing from data
		t.Run(name, func(t *testing.T) {
			actualSrv, actualErr := server.NewServer(tc.inputDB)

			// assert.EqualValues(t, tc.expectedSrv, actualSrv)
			assert.EqualValues(t, tc.expectedErr, actualErr)
			if tc.expectedErr == nil {
				assert.NotNil(t, actualSrv)
			}
		})
	}
}
