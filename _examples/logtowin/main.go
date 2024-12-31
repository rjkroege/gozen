package main

// go build

import (
	"flag"
	"log"
	"path/filepath"
	"os"
	"io"

	"github.com/rjkroege/gozen"
)

var filename = flag.String("f", "+Log", "the filename in the current directory to update")
var debug = flag.Bool("d",  false, "set for verbose debugging")

// need a wrapper for win to make it into a writable.


func main() {
	flag.Parse()
	if !*debug {
		log.SetOutput(io.Discard)
	}
	log.Println("hi", flag.Args())

	// have a single arg

	log.Println("filename", *filename)

	
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get current directory: %v", err)
	}

	targetfile := filepath.Join(cwd, *filename)

	win, err := gozen.Editinwin(targetfile)
	win.Clear()

	ww := gozen.NewWindowWriter("body", win)
	if _, err := io.Copy(ww, os.Stdin); err != nil {
		log.Fatalf("can't copy stdin to %s: %v", targetfile, err)
	}
}

