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
	router.POST("/callbackQuitGroupCommand", control.CallbackQuitGroup)
	router.POST("/callbackKillGroupCommand", control.CallbackKillGroupMember)
	router.POST("/callbackDisMissGroupCommand", control.CallbackDismissGroup)
	router.POST("/callbackBeforeJoinGroupCommand", control.CallbackApplyJoinGroupBefore)

	if err := engine.Run("192.168.62.1:8080"); err != nil {
		panic(err)
	}
}
