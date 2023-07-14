package main

import (
	"log"

	"github.com/LittleMoreInteresting/sql-tool/cmd"
)

//go:generate go run main.go sql -u root -p root --db bss_sys
//go:generate go run main.go file -f 1212
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
