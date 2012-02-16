// Browser core implementation. See also action.go, which implements
// user-initiated actions occuring outside of rendered content.

package main

var (
	exit     = make(chan interface{})
	actions  = make(chan Action)
	input    = make(chan interface{})
	requests = make(chan string) // Resource requests
)

func HandleActions() {
	for {
		select {
		case a := <-actions:
			a.Do()
		}
	}
}
