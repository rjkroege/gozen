package gozen

import (
	"9fans.net/go/acme"
)

// Pike-style options: [command center: Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
// I have written the simplest possible code that I need _now_. It is conceivable
// that I want something more sophisticated like https://golang.design/research/generic-option/#fn:1

type option func(*acme.Win) error

// Addtotag returns an option for Editinacme that adds the provided string
// to the Acme/Edwood tag.
func Addtotag(v string) option {
	return func(w *acme.Win) error {
		// capture v in a closure.
		return w.Fprintf("tag", v)
	}
}
