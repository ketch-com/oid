package main

import (
	"fmt"
	"go.ketch.com/lib/oid"
	"log"
	"os"
)

func main() {
	var o string
	var err error

	if len(os.Args) > 1 && os.Args[1] == "--random" {
		if o, err = oid.NewRandom(32); err != nil {
			log.Fatal(err)
		}
	} else if o, err = oid.NewOID(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(o)
}
