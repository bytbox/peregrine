package main

import (
	"os"

	"code.google.com/p/x-go-binding/xgb"
)

var (
	keymapRange = []byte{8, 255}

	keymap = make([][]xgb.Keysym, 256)
)

func GUIMain() {
	c, err := xgb.Dial(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}

	kmr, err := c.GetKeyboardMapping(keymapRange[0], keymapRange[1]-keymapRange[0])
	if err != nil { panic(err) }
	println(len(kmr.Keysyms))
	println(kmr.Length)
	j := 0
	for i := keymapRange[0]; i < keymapRange[1]; i++ {
		kss := make([]xgb.Keysym, kmr.KeysymsPerKeycode)
		for n := range kss {
			kss[n] = kmr.Keysyms[j]
			j++
		}
		keymap[i] = kss
	}

	screen := c.DefaultScreen()

	win := c.NewId()
	c.CreateWindow(0, win, screen.Root, 150, 150, 200, 200, 0, 0, 0, 0, nil)
	c.ChangeWindowAttributes(win, xgb.CWBackPixel|xgb.CWEventMask,
		[]uint32{
			screen.WhitePixel,
			xgb.EventMaskExposure |
			xgb.EventMaskKeyPress | xgb.EventMaskKeyRelease,
		})
	c.MapWindow(win)

	gc := c.NewId()
	c.CreateGC(
		gc, win,
		xgb.GCForeground | xgb.GCBackground,
		[]uint32{screen.BlackPixel, screen.WhitePixel},
	)

	for {
		reply, err := c.WaitForEvent()
		if err != nil { panic(err) }

		winGeo, err := c.GetGeometry(win)
		if err != nil { panic(err) }
		println(winGeo.Width)

		c.ClearArea(false, win, 0, 0, winGeo.Width, winGeo.Height)

		switch event := reply.(type) {
		case xgb.KeyPressEvent:
			println(keymap[event.Detail][0])
			points := make([]xgb.Point, 2)
			points[0].X = event.EventX
			points[0].Y = event.EventY
			c.PolyLine(xgb.CoordModeOrigin, win, gc, points)
		case xgb.KeyReleaseEvent:
		}
	}

	c.Close()
}
