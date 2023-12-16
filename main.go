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
	router.POST("/callbackGroupMsgReadCommand", control.CallbackGroupMsgReadCommand)
	router.POST("/callbackMsgModifyCommand", control.CallbackMsgModifyCommand)
	router.POST("/callbackAfterUpdateUserInfoCommand", control.CallbackAfterUpdateUserInfoCommand)
	router.POST("/callbackBeforeUserRegisterCommand", control.CallbackBeforeUserRegisterCommand)
	router.POST("/callbackAfterUserRegisterCommand", control.CallbackAfterUserRegisterCommand)
	router.POST("/callbackTransferGroupOwnerAfter", control.CallbackTransferGroupOwnerAfter)
	router.POST("/callbackBeforeSetFriendRemark", control.CallbackBeforeSetFriendRemark)
	router.POST("/callbackAfterSetFriendRemark", control.CallbackAfterSetFriendRemark)
	router.POST("/callbackSingleMsgRead", control.CallbackSingleMsgRead)
	router.POST("/callbackBeforeSendSingleMsgCommand", control.CallbackBeforeSendSingleMsgCommand)

	if err := engine.Run("0.0.0.0:18889"); err != nil {
		panic(err)
	}
}
