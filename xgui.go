package main

import (
	"os"

	"code.google.com/p/x-go-binding/xgb"
)

var (
	keymapRange = []byte{8, 255}

	keymap = make([][]xgb.Keysym, 256)
)

var (
	xC   *xgb.Conn
	xWin xgb.Id
	xGc  xgb.Id
	xGeo *xgb.GetGeometryReply

	lastMouse xgb.Point
)

func GUIInit() {
	c, err := xgb.Dial(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}

	kmr, err := c.GetKeyboardMapping(keymapRange[0], keymapRange[1]-keymapRange[0])
	if err != nil { panic(err) }
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

	xC, xWin, xGc = c, win, gc

	xGeo, err = xC.GetGeometry(win)
	if err != nil { panic(err) }
}

func GUIEventLoop() {
	for {
		reply, err := xC.WaitForEvent()
		if err != nil { panic(err) }

		switch event := reply.(type) {
		case xgb.KeyPressEvent:
			c, _ := getKeyChar(event.Detail, event.State)
			if c == 'q' && event.State & xgb.KeyButMaskControl != 0 {
				goto Exit
			}
		case xgb.KeyReleaseEvent:
		}
	}
Exit:
}

func getKeyChar(code byte, mask uint16) (rune, bool) {
	omask := mask
	mask &= 0x0003
	return rune(keymap[code][mask]), (omask & 0xffff-0x0003 == 0)
}

func GUIRender() {
	c, win, gc := xC, xWin, xGc

	c.ClearArea(true, win, 0, 0, xGeo.Width, xGeo.Height)

	c.PolyLine(xgb.CoordModeOrigin, win, gc, []xgb.Point{xgb.Point{0, 100}, xgb.Point{100,0}})
}
