package controller

import (
	"fmt"
	"net/http"
	"strings"
)

type LoginController struct {
}

func (c *LoginController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.RequestURI, "test1") {
		goto testGoto
	}
	fmt.Println("test")
	w.WriteHeader(200)
	w.Write([]byte("hello gweb!"))

testGoto:
	fmt.Println("testGoto")

	w.WriteHeader(200)
	w.Write([]byte("hello testGoto!"))
}
