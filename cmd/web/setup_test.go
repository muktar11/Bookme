//whatever is in setup_test.go will run before the other tests run use it as a supplier
package main 

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M){
//before start running other tests  do something inside this function and then run  test 
//and exit
	os.Exit(m.Run())
}

type myHandler struct{}
//object that satisfy the http handler interface
func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

}