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

func (con *Context) WriteString(str string) {
	WriterString(con.ResponseWriter, str)
}

func WriterString(w http.ResponseWriter, str string) {
	_, e := w.Write([]byte(str))
	if e != nil {
		log.Printf("返回前端数据出现异常：%s", e.Error())
	}
}
