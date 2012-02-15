package main

import (
	"os"

	"code.google.com/p/x-go-binding/xgb"
)

func GUIMain() {
	c, err := xgb.Dial(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}

	c.Close()
}
