package main

import (
	"gweb"
	_ "gweb/router"
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	log.Println("gweb start success ... ...")
	http.ListenAndServe(":8080", gweb.MRouter)
}
