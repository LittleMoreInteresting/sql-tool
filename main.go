package main

import (
	"log"

	"github.com/sql-tool/cmd"
)

//go:generate go run main.go struct -U root -P root --db bss_sys
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
