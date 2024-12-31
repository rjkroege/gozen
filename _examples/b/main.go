package main

// go build

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/rjkroege/gozen"
)

func main() {
	flag.Parse()
	log.Println("hi", flag.Args())

	// The argument to Editinacme needs to be an absolute path.
	ap, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatalf("can't abs %q: %v", flag.Arg(0), err)
	}

	if err := gozen.Editinacme(ap, gozen.Addtotag("hello, added with b"), gozen.Blinktag("")); err != nil {
		log.Fatalf("Editinacme failed for %s: %v", ap, err)
	}
}
