// Package logger contains all the logic needed to interact with a logger.
//
// The goal of this is to be a wrapper for an external logging library so that
// it becomes easier to swap it out if needed.
//
// If this is an unnecessary layer, this package can be discarded and the
// logging library can be used directly elsewhere.
//nolint:depguard // Allow the dependency here.
package logger

import (
	"errors"
	"log"
)

// Init carries on any initialisation needed by the logger of choice.
// Here is where you could pass a level argument and use it.
func Init() error {
	return errors.New("not implemented")
}

// Debugf prints a message on the debug level.
func Debugf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Infof prints a message on the info level.
func Infof(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Warningf prints a message on the info level.
func Warningf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Errorf prints a message on the info level.
func Errorf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}
