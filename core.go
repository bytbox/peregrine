// Browser core implementation. See also action.go, which implements
// user-initiated actions occuring outside of rendered content.

package main

import (
	"github.com/bytbox/waitmap.go"
)

var (
	exit     = make(chan interface{})
	actions  = make(chan Action)
	input    = make(chan interface{})
	requests = make(chan string) // Resource requests

	resources = waitmap.New()
)

func HandleActions() {
	for {
		select {
		case a := <-actions:
			a.Do()
		}
	}
}
