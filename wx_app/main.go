package main

import (
	routerInstance "wx_app/routes"
	// "encoding/json"
)


func main(){
	router:=routerInstance.InitRoutes()
	router.Run(":8888")
}