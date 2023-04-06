package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	gin.Recovery()
	_ = engine.SetTrustedProxies(nil)
	if e := InitApp(engine); e != nil { // register urls of all app
		log.Fatalf("Application register failed, err = %v", e)
		return
	}
	if e := engine.Run("127.0.0.1:8080"); e != nil {
		log.Fatalf("Gin engine run failed, err = %v", e)
		return
	}
}
