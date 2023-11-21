package model

type CallbackCommand string

type CreateGroupReq struct {
	OperationID            string               `json:"operationID,omitempty"`
	CallbackCommand        string               `json:"callbackCommand,omitempty"`
	GroupID                string               `json:"groupID,omitempty"`
	GroupName              string               `json:"groupName,omitempty"`
	Notification           string               ` json:"notification,omitempty"`
	Introduction           string               ` json:"introduction,omitempty"`
	FaceURL                string               `json:"faceURL,omitempty"`
	OwnerUserID            string               ` json:"ownerUserID,omitempty"`
	CreateTime             int64                `json:"createTime,omitempty"`
	MemberCount            uint32               `json:"memberCount,omitempty"`
	Ex                     string               `json:"ex,omitempty"`
	Status                 int32                ` json:"status,omitempty"`
	CreatorUserID          string               ` json:"creatorUserID,omitempty"`
	GroupType              int32                `json:"groupType,omitempty"`
	NeedVerification       int32                `json:"needVerification,omitempty"`
	LookMemberInfo         int32                ` json:"lookMemberInfo,omitempty"`
	ApplyMemberFriend      int32                `json:"applyMemberFriend,omitempty"`
	NotificationUpdateTime int64                ` json:"notificationUpdateTime,omitempty"`
	NotificationUserID     string               `json:"notificationUserID,omitempty"`
	GroupAddMemberInfo     []GroupAddMemberInfo ` json:"groupAddMemberInfo,omitempty"`
}

type GroupAddMemberInfo struct {
	UserID    string `json:"userID,omitempty"`
	RoleLevel int32  `json:"roleLevel,omitempty"`
}

type CommonCallbackResp struct {
	ActionCode int32  `json:"actionCode,omitempty"`
	ErrCode    int32  ` json:"errCode,omitempty"`
	ErrMsg     string ` json:"errMsg,omitempty"`
	ErrDlt     string `json:"errDlt,omitempty"`
	NextCode   int32  `json:"nextCode,omitempty"`
}

type CreateGroupResp struct {
	CommonCallbackResp
	GroupID           *string `json:"groupID"`
	GroupName         *string `json:"groupName"`
	Notification      *string `json:"notification"`
	Introduction      *string `json:"introduction"`
	FaceURL           *string `json:"faceURL"`
	OwnerUserID       *string `json:"ownerUserID"`
	Ex                *string `json:"ex"`
	Status            *int32  `json:"status"`
	CreatorUserID     *string `json:"creatorUserID"`
	GroupType         *int32  `json:"groupType"`
	NeedVerification  *int32  `json:"needVerification"`
	LookMemberInfo    *int32  `json:"lookMemberInfo"`
	ApplyMemberFriend *int32  `json:"applyMemberFriend"`
}

type CreateGroupAfterReq struct {
	OperationID            string               `json:"operationID,omitempty"`
	CallbackCommand        string               `json:"callbackCommand,omitempty"`
	GroupID                string               `json:"groupID,omitempty"`
	GroupName              string               `json:"groupName,omitempty"`
	Notification           string               ` json:"notification,omitempty"`
	Introduction           string               ` json:"introduction,omitempty"`
	FaceURL                string               `json:"faceURL,omitempty"`
	OwnerUserID            string               ` json:"ownerUserID,omitempty"`
	CreateTime             int64                `json:"createTime,omitempty"`
	MemberCount            uint32               `json:"memberCount,omitempty"`
	Ex                     string               `json:"ex,omitempty"`
	Status                 int32                ` json:"status,omitempty"`
	CreatorUserID          string               ` json:"creatorUserID,omitempty"`
	GroupType              int32                `json:"groupType,omitempty"`
	NeedVerification       int32                `json:"needVerification,omitempty"`
	LookMemberInfo         int32                ` json:"lookMemberInfo,omitempty"`
	ApplyMemberFriend      int32                `json:"applyMemberFriend,omitempty"`
	NotificationUpdateTime int64                ` json:"notificationUpdateTime,omitempty"`
	NotificationUserID     string               `json:"notificationUserID,omitempty"`
	GroupAddMemberInfo     []GroupAddMemberInfo ` json:"groupAddMemberInfo,omitempty"`
}

type CallbackQuitGroupReq struct {
	CallbackCommand `json:"callbackCommand"`
	GroupID         string `json:"groupID"`
	UserID          string `json:"userID"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackQuitGroupResp struct {
	CommonCallbackResp
}

type CallbackKillGroupMemberReq struct {
	CallbackCommand `json:"callbackCommand"`
	GroupID         string   `json:"groupID"`
	KickedUserIDs   []string `json:"kickedUserIDs"`
	Reason          string   `json:"reason"`
	EventTime       int64    `json:"eventTime"`
}

type CallbackKillGroupMemberResp struct {
	CommonCallbackResp
}

type CallbackDisMissGroupReq struct {
	CallbackCommand `json:"callbackCommand"`
	GroupID         string   `json:"groupID"`
	OwnerID         string   `json:"ownerID"`
	GroupType       string   `json:"groupType"`
	MembersID       []string `json:"membersID"`
	EventTime       int64    `json:"eventTime"`
}

type CallbackDisMissGroupResp struct {
	CommonCallbackResp
}

type CallbackJoinGroupReq struct {
	CallbackCommand `json:"callbackCommand"`
	GroupID         string `json:"groupID"`
	GroupType       string `json:"groupType"`
	ApplyID         string `json:"applyID"`
	ReqMessage      string `json:"reqMessage"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackJoinGroupResp struct {
	CommonCallbackResp
}
