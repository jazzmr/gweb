package context

import (
	"net/http"
)

type RequestUri struct {
	ContextPath   string
	Mapping       string
	Method        string
	PathParams    []string
	RequestParams map[string]string
}

type Context struct {
	ResponseWriter http.ResponseWriter
	RequestUri     *RequestUri
	Request        *http.Request
}

func (con *Context) WriteString(str string) {
	con.ResponseWriter.Write([]byte(str))
}
