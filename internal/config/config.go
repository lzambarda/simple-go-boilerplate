// Package config contains all the code needed to correctly read and set up
// configuration in the repo.
//
// Here libraries such as https://github.com/caarlos0/env or
// github.com/spf13/viper can be used.
package config

import (
	"errors"
	"os"
)

const (
	fooName = "FOO"
)

// Foo is the variable that does the fooing.
var Foo string

// LoadEnv loads both environment variables and flags. This is not done inside
// an init function so that we can control when things are initialised.
//
// Only variables GLOBAL to the entire repo should be defined here, otherwise a
// Config-based approach is preferred to avoid having to remember what
// environment variables are needed by each of the entrypoints inside cmd.
func LoadEnv() error {
	Foo = os.Getenv(fooName)

	if Foo == "" {
		return errors.New("the variable FOO is not set") //nolint:err113 // Fine here.
	}

	return nil
}
