package main

import (
	routerInstance "routes"
	// "encoding/json"
)


func main(){
	router:=routerInstance.InitRoutes()
	router.Run(":8888")
}