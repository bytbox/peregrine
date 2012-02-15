package main

type Config struct {
	Homepage string
}

var config Config

func init() {
	// TODO actually read config
	config.Homepage = "about:blank"
}
