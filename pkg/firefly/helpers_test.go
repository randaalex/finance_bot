package firefly_test

import (
	"bytes"
	"io"
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func ReadCloserToString(input io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)

  	return buf.String()
}
