package main

var (
	exit    chan interface{}
	actions chan Action
	input   chan interface{}
)

func init() {
	exit = make(chan interface{})
	actions = make(chan Action)
	input = make(chan interface{})
}

func HandleActions() {
	for {
		select {
		case a := <-actions:
			a.Do()
		}
	}
}
