package controller

import "net/http"

type HelloController struct {
	Url    string
	Method string
}

func (c *HelloController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (c *HelloController) hello1() {

}
