package main

import (
	"github.com/magiskboy/katalog/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
