package main 


import (
	"fmt"
	"github.com/go-chi/chi"
	"Bookme/internal/config"
	"testing"
)

func TestRoutes(t *testing.T){
	//need an app config and return handler
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type){
	//routes return mux which is a pointer of CHI MUX
	case *chi.Mux: 
		//do nothingg test passed 
	default: 
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}