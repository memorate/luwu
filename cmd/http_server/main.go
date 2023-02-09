package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	gin.Recovery()
	app, err := InitializeApplication()
	if err != nil {
		fmt.Printf("InitializeApplication failed, err = %v", err)
		return
	}
	if e := app.Register(engine); e != nil { // register urls of all app
		fmt.Printf("Application register failed, err = %v", e)
		return
	}
	if e := engine.Run("127.0.0.1:8080"); e != nil {
		fmt.Printf("Gin engine run failed, err = %v", e)
		return
	}
}
