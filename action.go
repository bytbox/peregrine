package main

type Action interface {
	Do()
}

// Simple action kinds.
const (
	ExitAction = iota
)

type SimpleAction struct {
	kind int
}

func Simple(kind int) SimpleAction {
	return SimpleAction{kind}
}

func (a SimpleAction) Do() {
	switch a.kind {
	case ExitAction:
		exit <- nil
	default:
		panic("Unrecognized kind")
	}
}

type NavigateAction struct {
	dest string
}

func Navigate(dest string) NavigateAction {
	return NavigateAction{dest}
}

func (a NavigateAction) Do() {

}
