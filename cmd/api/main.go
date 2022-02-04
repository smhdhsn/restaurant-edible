package main

import (
	"fmt"
	"log"

	"github.com/smhdhsn/food/internal/config"
)

// main is the main application entry.
func main() {
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
}
