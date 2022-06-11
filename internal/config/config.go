// Package config contains all the code needed to correctly read and set up
// configuration in the repo.
//
// Here libraries such as github.com/spf13/viper can be used.
package config

import "os"

const (
	fooName = "FOO"
)

var (
	Foo string
)

func init() {
	Foo = os.Getenv(fooName)
}
