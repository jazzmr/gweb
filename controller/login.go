package controller

import (
	"net/http"
)

type LoginController struct {
}

func (c *LoginController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello gweb!"))
}
