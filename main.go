package main

import (
	"log"
	"strings"
)

func main() {
	log.Printf("%s actions> HELLO, FROM MAIN FUNCTION %s",
		strings.Repeat("=", 16),
		strings.Repeat("=", 16),
	)
}
