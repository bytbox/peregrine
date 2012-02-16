package main

// Possible values for widgetPosition.rel
const (
	WP_UL = iota
	WP_LL
	WP_UR
	WP_LR
)

type WidgetPosition struct {
	x, y int
	rel  int
}

func (p WidgetPosition) Real(sx, sy int) (x int, y int) {
	switch p.rel {
	case WP_UL:
		x, y = p.x, p.y
	case WP_LL:
		x, y = p.x, sy-p.y
	case WP_UR:
		x, y = sx-p.x, p.y
	case WP_LR:
		x, y = sx-p.x, sy-p.y
	default:
		panic("Bad WidgetPosition.rel")
	}
	return
}
