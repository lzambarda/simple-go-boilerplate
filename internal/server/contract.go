package server

// Database is the contract used by the server. This should contain all the
// methods that the server intends to use.
// This is a good entrypoint for mock generating libraries such as:
// github.com/vektra/mockery
type Database interface {
	// DoSomething does what is written on the tin.
	DoSomething() (res string, err error)
}
