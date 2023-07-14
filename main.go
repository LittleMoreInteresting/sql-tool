package main

import (
	"log"

	"github.com/LittleMoreInteresting/sql-tool/cmd"
)

//go:generate go run main.go struct -u root -p root --db bss_sys
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
