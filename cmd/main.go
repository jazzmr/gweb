package main

import (
	"fmt"
	"gweb"
	"gweb/conf"
	_ "gweb/router"
	"log"
	"net/http"
	"time"
)

func main() {

	c, err := conf.GetConfig()

	if err != nil {
		log.Fatalf("get config err, err is %v", err)
	}

	time.Sleep(1 * time.Second)

	fmt.Println(c.Server)

	log.Println("gweb start success ... ...")
	http.ListenAndServe(fmt.Sprintf(":%d", c.Server.Port), gweb.MRouter)
}
