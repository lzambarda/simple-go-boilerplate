// Package main is the entrypoint of the program.
//
// Having a single entrypoing means that you will only get one binary which does
// everything.
// You can create several entrypoints by creating different folder in here and
// have a package main in each one:
//
//  cmd
//  ├── foo
//  │   ├── main.go
//  │   └── foo.go
//  └── bar
//      └── main.go
//
// With the above layout you can then build separate binaries such as:
//
// go build -o bin/foo cmd/foo/*.go
// go build -o bin/bar cmd/bar/*.go
package main
