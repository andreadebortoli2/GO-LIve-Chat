package main

import (
	"net/http"
	"testing"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/config"
)

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func TestRouter(t *testing.T) {
	var app config.AppConfig

	r := Router(&app)

	switch v := r.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not an http.handler, but is %T", v)
	}
}

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing test passed
	default:
		t.Errorf("type is not an http.Handler, but is %T", v)
	}
}

func TestLoadSessio(t *testing.T) {
	var myH myHandler

	h := LoadSession(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not an http.handler, but is %T", v)
	}
}

func TestAuthMiddleware(t *testing.T) {
	var myH myHandler

	a := authMiddleware(&myH)

	switch v := a.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not an http.Handler, but is %T", v)

	}
}
func TestAdminAuthMiddleware(t *testing.T) {
	var myH myHandler

	a := adminAuthMiddleware(&myH)

	switch v := a.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not an http.Handler, but is %T", v)

	}
}
