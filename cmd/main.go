package main

import (
	"fmt"

	"github.com/ORGNAME/REPONAME/internal/config"
	"github.com/ORGNAME/REPONAME/internal/logger"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("Got an env variable: %q\n", config.Foo)
	logger.Infof("I can even use a logger package!")
}
