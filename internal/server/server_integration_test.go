//go:build integration

package server_test

import (
	"embed"
	"testing"
)

// Embedding testdata makes it easier to write tests without having to worry
// about where they are called from.
//
//go:embed testdata
var testdata embed.FS

// TestServerIntegration is a simple integration test.
func TestServerIntegration(t *testing.T) {
	t.Log("this will only run with make integration_test")

	b, err := testdata.ReadFile("testdata/fixture1.txt")
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}

	t.Logf("contents of fixture: %s", string(b))
}
