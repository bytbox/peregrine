// Functions and resources that will be shared by all GUI implementations.

/*

The following objects define the singleton GUI interface, and must be
implemented outside of this file:

	- GUIInit      func()
	- GUIEventLoop func()

The following objects are provided by this file as supplements:

	- GUIMain       func()
	- GUIRenderLoop func()

*/

package main

import (
	"time"
)

const (
	GUI_FMS = 30. // ms per frame
)

type Painter interface {
}

func GUIMain() {
	GUIInit()
	go GUIRenderLoop()
	GUIEventLoop()
	exit <- nil
}

func GUIRenderLoop() {
	ticker := time.Tick(1e6*GUI_FMS)
	for {
		GUIRender()
		<-ticker
	}
}
