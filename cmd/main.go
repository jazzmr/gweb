package main

import (
	"fmt"
	"gweb"
	"gweb/conf"
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

	h := &gweb.Controller{
		Handler: http.HandlerFunc(myFunc),
		Pattern: "localhost",
	}

	log.Println("gweb start success ... ...")
	http.ListenAndServe(fmt.Sprintf(":%d", c.Server.Port), h)
}

func myFunc(rw http.ResponseWriter, r *http.Request) {

	time.Sleep(5 * time.Second)

	fmt.Println("hello world!")
	rw.Write([]byte("hello world!"))
}
