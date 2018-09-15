package context

import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
}

func (con *Context) WriteString(str string) {
	con.ResponseWriter.Write([]byte(str))
}
