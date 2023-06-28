package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/:lang/var", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/:lang/var"), []string{":lang", "var"})
	if !ok {
		t.Fatal("test fail")
	}
}

func TestGetRouter(t *testing.T) {
	r := newTestRouter()
	_, ps := r.getRouter("GET", "/go/var")
	fmt.Println(ps)
}
