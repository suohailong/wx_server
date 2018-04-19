package main

import (
	// routerInstance "./routes"
	// "encoding/json"
	routerInstance "wx_app/routes"
)


func main(){
	router:=routerInstance.InitRoutes()
	router.Run(":8888")
}