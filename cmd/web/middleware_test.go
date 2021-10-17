package main

import (
	"testing"
	"net/http"
	"fmt"
)


func TestNoSurf(t *testing.T){
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing

	default:
		t.Error(fmt.Sprintf("type is not heep.Handler, but is %T", v))	
		
	}
}


func TestSessionLoad(t *testing.T){
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing

	default:
		t.Error(fmt.Sprintf("type is not heep.Handler, but is %T", v))	
		
	}
}