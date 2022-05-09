package mian

import (
	"log"

	"github.com/sql-tool/cmd"
)



func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}