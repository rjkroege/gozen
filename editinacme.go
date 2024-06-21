package gozen

import (
	"fmt"
	"log"
	"strings"

	"9fans.net/go/acme"
)

// Editinacme directly opens plumbstring in Acme/Edwood because regular
// plumb can't handle the paths found in the Go package database.
// Note that paths in plumbstring need to be absolute.
func Editinacme(plumbstring string) error {
	chunks := strings.Split(plumbstring, ":")
	if len(chunks) > 2 {
		return fmt.Errorf("plumbhelper bad plumb address string")
	}
	fn := chunks[0]
	addr := ""
	if len(chunks) > 1 {
		addr = chunks[1]
	}
	log.Println("plumbhelper", fn, addr)

	// Two choices: we already have the Window open.
	wins, err := acme.Windows()
	if err != nil {
		return fmt.Errorf("plumbhelper acme.Windows")
	}

	win := (*acme.Win)(nil)
	for _, wi := range wins {
		// log.Println("wi", wi.Name)
		if wi.Name == fn {
			win, err = acme.Open(wi.ID, nil)
			if err != nil {
				return fmt.Errorf("plumbhelper acme.Open")
			}
			break
		}
	}

	if win == nil {
		log.Println("plumbhelper making a new window")
		var err error
		win, err = acme.New()
		if err != nil {
			return fmt.Errorf("plumbhelper acme.New: %v", err)
		}

		if err := win.Ctl("nomark"); err != nil {
			return fmt.Errorf("plumbhelper win.Ctl nomark: %v", err)
		}

		if err := win.Name(fn); err != nil {
			return fmt.Errorf("plumbhelper win.Name: %v", err)
		}

		if err := win.Ctl("get"); err != nil {
			return fmt.Errorf("plumbhelper win.Ctl get: %v", err)
		}

		if err = win.Ctl("mark"); err != nil {
			return fmt.Errorf("plumbhelper %q: %v", "mark", err)
		}

		if err = win.Ctl("clean"); err != nil {
			return fmt.Errorf("plumbhelper %q: %v", "clean", err)
		}
	}

	if err := win.Addr(string(addr)); err != nil {
		return fmt.Errorf("plumbhelper win.Addr: %v", err)
	}
	if err := win.Ctl("dot=addr\nshow\n"); err != nil {
		return fmt.Errorf("plumbhelper win.Addr: %v", err)
	}

	return nil
}
