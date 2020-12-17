package main

import (
	"fmt"
	"go.ketch.com/lib/oid"
	"log"
)

func main() {
	o, err := oid.NewOID()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(o)
}
