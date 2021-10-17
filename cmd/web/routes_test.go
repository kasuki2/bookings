package main

import (
	"testing"
"github.com/go-chi/chi/v5"
"github.com/kasuki2/bookings/internal/config"
"fmt"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing 
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but type is %T", v))	
	}
}
