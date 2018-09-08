package gweb

import (
	"fmt"
	"net/http"
)

type Controller struct {
	Handler http.Handler
	Pattern string
}

func (c *Controller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Controller ServeHTTP... ...")

	c.Handler.ServeHTTP(rw, r)

}
