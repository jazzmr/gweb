package context

import (
	"log"
	"net/http"
)

type RequestUri struct {
	ContextPath string
	Mapping     string
}

type Context struct {
	ResponseWriter http.ResponseWriter
	RequestUri     *RequestUri
	Request        *http.Request
}

func (ctx *Context) Reset(w http.ResponseWriter, r *http.Request, uri *RequestUri) {
	ctx.Request = r
	ctx.ResponseWriter = w
	ctx.RequestUri = uri
}

func (ctx *Context) WriteString(str string) {
	WriterString(ctx.ResponseWriter, str)
}

/**
create a new empty context
*/
func NewContext() *Context {
	return &Context{}
}

func WriterString(w http.ResponseWriter, str string) {
	_, e := w.Write([]byte(str))
	if e != nil {
		log.Printf("返回前端数据出现异常：%s", e.Error())
	}
}
