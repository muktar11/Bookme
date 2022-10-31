//testing the type of http coming and going
package main

import (
	"fmt"
	"net/http"
	"testing"
)

//test the middleware NoSurf which returns a http handler\
// we use setup_test to supply next http handler
func TestNoSurf(t *testing.T){
	//appluy next http handler here
	
	var myH myHandler

	h := NoSurf(&myH) 

	// assing the type of h to value of v
	switch v := h.(type){
	case http.Handler:
		//if the type is http handler do nothing
	default:
		t.Error(fmt.Sprintf("type is not http Handler, but is %T", v))
	}
}

//test SessionLoad which take http handler and return handler

func TestSessionLoad(t *testing.T){
	//appluy next http handler here
	
	var myH myHandler

	h := SessionLoad(&myH) 

	// assing the type of h to value of v
	switch v := h.(type){
	case http.Handler:
		//if the type is http handler do nothing
	default:
		t.Error(fmt.Sprintf("type is not http Handler, but is %T", v))
	}
}