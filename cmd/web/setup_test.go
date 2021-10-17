package main

import (
	"testing"
	"os"
	"net/http"
)

func TestMain(m *testing.M) {



	os.Exit(m.Run())

}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

}