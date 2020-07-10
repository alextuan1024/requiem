package main

import (
	"github.com/alextuan1024/requiem/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
