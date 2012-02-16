package main

import (
	"errors"
	"io/ioutil"
	"net/url"
)

func Fetcher() {
	for {
		select {
		case req := <-requests:
			res, err := Access(req)
			if err != nil {
				panic(err)
			}
			println(string(res.Data))
		}
	}
}

type Resource struct {
	Data []byte
	Meta map[string]string
}

func RawResource(bs []byte) *Resource {
	return &Resource{bs, make(map[string]string)}
}

// Accessors are expected to be thread-safe.
type Accessor interface {
	Access(string) (*Resource, error)
}

// An HTTPAccessor can be used as a regular accessor, but can also make more
// elaborate requests as triggered by forms, hyperlinks, js, and the like.
type HTTPAccessor interface {
	Accessor
}

type BasicAccessor struct {
	Get func(string) (*Resource, error)
}

func (b BasicAccessor) Access(url string) (*Resource, error) {
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

func Access(urlstr string) (*Resource, error) {
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

func aboutAccess(urlstr string) (*Resource, error) {
	_, err := url.Parse(urlstr)
	if err != nil { return nil, err }
	return RawResource([]byte("Not yet implemented.")), nil
}

func fileAccess(urlstr string) (*Resource, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(u.Path)
	return RawResource(bs), err
}
