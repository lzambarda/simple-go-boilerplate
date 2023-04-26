// Package config contains all the code needed to correctly read and set up
// configuration in the repo.
//
// Here libraries such as github.com/spf13/viper can be used.
package config

import "os"

const (
	fooName = "FOO"
)

var Foo string

// LoadEnv loads both environment variables and flags. This is not done inside
// an init function so that we can control when things are initialised.
func LoadEnv() error {
	Foo = os.Getenv(fooName)

	return nil
}
