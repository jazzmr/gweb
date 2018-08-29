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

	//  静态服务器
	//router := mux.NewRouter().StrictSlash(true)
	//router.Handle("/", http.FileServer(http.Dir("controller"))).Methods("GET")
	////http.Handle("/", http.FileServer(http.Dir("controller")))
	//http.ListenAndServe(":8080", router)

	c, err := conf.GetConfig()

	if err != nil {
		log.Fatalf("get config err, err is %v", err)
	}

	time.Sleep(1 * time.Second)

	fmt.Println(c.Server)

	log.Println("gweb start success ... ...")
	http.ListenAndServe(fmt.Sprintf(":%d", c.Server.Port), gweb.GRouter)
}
