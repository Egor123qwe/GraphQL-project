package main

import (
	"github.com/Egor123qwe/GraphQL-project/internal/app"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
