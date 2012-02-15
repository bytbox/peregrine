package main

import (
	"errors"
	"io/ioutil"
	"net/url"
)

// Accessors are expected to be thread-safe.
type Accessor interface {
	Access(string) ([]byte, error)
}

// An HTTPAccessor can be used as a regular accessor, but can also make more
// elaborate requests as triggered by forms, hyperlinks, js, and the like.
type HTTPAccessor interface {
	Accessor
}

type BasicAccessor struct {
	Get func(string) ([]byte, error)
}

func (b BasicAccessor) Access(url string) ([]byte, error) {
	return b.Get(url)
}

// The accessor protocol map associates protocol names ('http', 'file', etc.)
// with appropriate accessor objects.
var Accessors map[string]Accessor

func init() {
	Accessors = map[string]Accessor{
		"about": BasicAccessor{aboutAccess},
		"file":  BasicAccessor{fileAccess},
	}
}

func Access(urlstr string) ([]byte, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	a, ok := Accessors[u.Scheme]
	if !ok {
		return nil, errors.New("invalid scheme")
	}
	return a.Access(urlstr)
}

func aboutAccess(urlstr string) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func fileAccess(urlstr string) ([]byte, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(u.Path)
}
