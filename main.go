package main

import (
	"log"

	"github.com/sql-tool/cmd"
)

//go:generate go run main.go sql struct -U root -P root --db blog_service --table blog_tag
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}