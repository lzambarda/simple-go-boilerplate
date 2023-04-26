package main

import (
	"fmt"

	"github.com/ORGNAME/REPONAME/internal/config"
	"github.com/ORGNAME/REPONAME/internal/logger"
)

func main() {
	fmt.Println("Hello world!")
	logger.Infof("I can even use a logger package!")
	err := config.LoadEnv()
	if err != nil {
		logger.Fatalf("load env: %s", err)
	}
	fmt.Printf("Got an env variable: %q\n", config.Foo)
}
