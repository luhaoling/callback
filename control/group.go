package control

import (
	"call-back-http/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "ping")
}

func CreateGroupBefore(c *gin.Context) {
	var req model.CreateGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	respErr := model.CommonCallbackResp{
		ActionCode: 0,
		ErrCode:    200,
		ErrMsg:     "Success",
		ErrDlt:     "Successful",
		NextCode:   2,
	}

	resp := &model.CreateGroupResp{
		CommonCallbackResp: respErr,
		GroupID:            req.GroupID,
		GroupName:          req.GroupName,
		Notification:       req.Notification,
		Introduction:       req.Introduction,
		FaceURL:            req.FaceURL,
		OwnerUserID:        req.OwnerUserID,
		Ex:                 req.Ex,
		Status:             req.Status,
		CreatorUserID:      req.CreatorUserID,
		GroupType:          req.GroupType,
		NeedVerification:   req.NeedVerification,
		LookMemberInfo:     req.LookMemberInfo,
		ApplyMemberFriend:  req.ApplyMemberFriend,
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CreateGroupAfter(c *gin.Context) {
	var req model.CreateGroupAfterReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CommonCallbackResp{
		ActionCode: 1,
		ErrCode:    200,
		ErrMsg:     "Success",
		ErrDlt:     "Successful",
		NextCode:   2,
	}

	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackQuitGroup(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CommonCallbackResp{
		ActionCode: 0,
		ErrCode:    200,
		ErrMsg:     "Success",
		ErrDlt:     "Successful",
		NextCode:   2,
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackKillGroupMember(c *gin.Context) {
	var req model.CallbackKillGroupMemberReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)

	resp := &model.CallbackKillGroupMemberResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackDismissGroup(c *gin.Context) {
	var req model.CallbackDisMissGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackDisMissGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackApplyJoinGroupBefore(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackJoinGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackQuitGroupCommand(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackJoinGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackKillGroupCommand(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackJoinGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackDisMissGroupCommand(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackJoinGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackBeforeJoinGroupCommand(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackJoinGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackGroupMsgReadCommand(c *gin.Context) {
	var req model.CallbackGroupMsgReadReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackGroupMsgReadResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

// 还需补充
func CallbackMsgModifyCommand(c *gin.Context) {
	var req model.CallbackMsgModifyCommandReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackMsgModifyCommandResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackAfterUpdateUserInfoCommand(c *gin.Context) {
	var req model.CallbackJoinGroupReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackJoinGroupResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackBeforeUserRegisterCommand(c *gin.Context) {
	var req model.CallbackBeforeUserRegisterReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackBeforeUserRegisterResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
		Users: req.Users,
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackAfterUserRegisterCommand(c *gin.Context) {
	var req model.CallbackAfterUserRegisterReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackAfterUserRegisterResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackTransferGroupOwnerAfter(c *gin.Context) {
	var req model.CallbackTransferGroupOwnerReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackTransferGroupOwnerResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackBeforeSetFriendRemark(c *gin.Context) {
	var req model.CallbackBeforeSetFriendRemarkReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackBeforeSetFriendRemarkResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackAfterSetFriendRemark(c *gin.Context) {
	var req model.CallbackAfterSetFriendRemarkReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackAfterSetFriendRemarkResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}

func CallbackSingleMsgRead(c *gin.Context) {
	var req model.CallbackSingleMsgReadReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Printf("%#v\n", req)
	resp := &model.CallbackSingleMsgReadResp{
		CommonCallbackResp: model.CommonCallbackResp{
			ActionCode: 0,
			ErrCode:    200,
			ErrMsg:     "Success",
			ErrDlt:     "Successful",
			NextCode:   2,
		},
	}
	fmt.Printf("%#v", resp)
	c.JSON(http.StatusOK, resp)
}
