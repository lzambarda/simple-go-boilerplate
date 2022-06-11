//go:build integration

package server_test

import "testing"

// TestServerIntegration is a simple integration test
func TestServerIntegration(t *testing.T) {
	t.Log("this will only run with make integration_test")
}
