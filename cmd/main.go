package main

import "gweb"
import _ "gweb/router"

func main() {

	//  静态服务器
	//router := mux.NewRouter().StrictSlash(true)
	//router.Handle("/", http.FileServer(http.Dir("controller"))).Methods("GET")
	////http.Handle("/", http.FileServer(http.Dir("controller")))
	//http.ListenAndServe(":8080", router)

	gweb.Run()
}
