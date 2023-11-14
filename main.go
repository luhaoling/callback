package main

import (
	"call-back-http/control"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router := engine.Group("/sdkName")
	router.POST("/callbackBeforeCreateGroupCommand", control.CreateGroupBefore)
	router.POST("/callbackAfterCreateGroupCommand", control.CreateGroupAfter)
	if err := engine.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
