package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = `0.0.1`
)

var (
	version = flag.Bool("version", false, "display version information")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [url]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if *version {
		fmt.Printf("peregrine %s\n", VERSION)
	}

	args := flag.Args()
	var url string
	switch len(args) {
	case 0:
		url = config.Homepage
	case 1:
		url = args[0]
	default:
		flag.Usage()
		return
	}

	go GUIMain()
	go Fetcher()
	go HandleActions()

	actions <- Navigate(url)

	<-exit
}
