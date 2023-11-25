package model

import "google.golang.org/protobuf/runtime/protoimpl"

type CallbackCommand string

type CommonCallbackReq struct {
	SendID           string   `json:"sendID"`
	CallbackCommand  string   `json:"callbackCommand"`
	ServerMsgID      string   `json:"serverMsgID"`
	ClientMsgID      string   `json:"clientMsgID"`
	OperationID      string   `json:"operationID"`
	SenderPlatformID int32    `json:"senderPlatformID"`
	SenderNickname   string   `json:"senderNickname"`
	SessionType      int32    `json:"sessionType"`
	MsgFrom          int32    `json:"msgFrom"`
	ContentType      int32    `json:"contentType"`
	Status           int32    `json:"status"`
	CreateTime       int64    `json:"createTime"`
	Content          string   `json:"content"`
	Seq              uint32   `json:"seq"`
	AtUserIDList     []string `json:"atUserList"`
	SenderFaceURL    string   `json:"faceURL"`
	Ex               string   `json:"ex"`
}

type OfflinePushInfo struct {
	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	Desc          string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc"`
	Ex            string `protobuf:"bytes,3,opt,name=ex,proto3" json:"ex"`
	IOSPushSound  string `protobuf:"bytes,4,opt,name=iOSPushSound,proto3" json:"iOSPushSound"`
	IOSBadgeCount bool   `protobuf:"varint,5,opt,name=iOSBadgeCount,proto3" json:"iOSBadgeCount"`
	SignalInfo    string `protobuf:"bytes,6,opt,name=signalInfo,proto3" json:"signalInfo"`
}

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
	GroupID           string `json:"groupID"`
	GroupName         string `json:"groupName"`
	Notification      string `json:"notification"`
	Introduction      string `json:"introduction"`
	FaceURL           string `json:"faceURL"`
	OwnerUserID       string `json:"ownerUserID"`
	Ex                string `json:"ex"`
	Status            int32  `json:"status"`
	CreatorUserID     string `json:"creatorUserID"`
	GroupType         int32  `json:"groupType"`
	NeedVerification  int32  `json:"needVerification"`
	LookMemberInfo    int32  `json:"lookMemberInfo"`
	ApplyMemberFriend int32  `json:"applyMemberFriend"`
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

type CallbackBeforeSetFriendRemarkReq struct {
	CallbackCommand `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	OwnerUserID     string `json:"ownerUserID"`
	FriendUserID    string `json:"friendUserID"`
	Remark          string `json:"remark"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackBeforeSetFriendRemarkResp struct {
	CommonCallbackResp
	Remark string `json:"remark"`
}

type CallbackAfterSetFriendRemarkReq struct {
	CallbackCommand `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	OwnerUserID     string `json:"ownerUserID"`
	FriendUserID    string `json:"friendUserID"`
	Remark          string `json:"remark"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackAfterSetFriendRemarkResp struct {
	CommonCallbackResp
}

type CallbackAfterSetGroupMemberInfoReq struct {
	CallbackCommand `json:"callbackCommand"`
	GroupID         string  `json:"groupID"`
	UserID          string  `json:"userID"`
	Nickname        *string `json:"nickName"`
	FaceURL         *string `json:"faceURL"`
	RoleLevel       *int32  `json:"roleLevel"`
	Ex              *string `json:"ex"`
	EventTime       int64   `json:"eventTime"`
}

type CallbackAfterSetGroupMemberInfoResp struct {
	CommonCallbackResp
}

type CallbackTransferGroupOwnerReq struct {
	CallbackCommand `json:"callbackCommand"`
	GroupID         string `json:"groupID"`
	OldOwnerUserID  string `json:"oldOwnerUserID"`
	NewOwnerUserID  string `json:"newOwnerUserID"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackTransferGroupOwnerResp struct {
	CommonCallbackResp
}

type CallbackGroupMsgReadReq struct {
	CallbackCommand `json:"callbackCommand"`
	SendID          string `json:"sendID"`
	ReceiveID       string `json:"receiveID"`
	UnreadMsgNum    int64  `json:"unreadMsgNum"`
	ContentType     int64  `json:"contentType"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackGroupMsgReadResp struct {
	CommonCallbackResp
}

type CallbackMsgModifyCommandReq struct {
	CommonCallbackReq
}

type CallbackMsgModifyCommandResp struct {
	CommonCallbackResp

	Content          *string          `json:"content"`
	RecvID           *string          `json:"recvID"`
	GroupID          *string          `json:"groupID"`
	ClientMsgID      *string          `json:"clientMsgID"`
	ServerMsgID      *string          `json:"serverMsgID"`
	SenderPlatformID *int32           `json:"senderPlatformID"`
	SenderNickname   *string          `json:"senderNickname"`
	SenderFaceURL    *string          `json:"senderFaceURL"`
	SessionType      *int32           `json:"sessionType"`
	MsgFrom          *int32           `json:"msgFrom"`
	ContentType      *int32           `json:"contentType"`
	Status           *int32           `json:"status"`
	Options          *map[string]bool `json:"options"`
	OfflinePushInfo  *OfflinePushInfo `json:"offlinePushInfo"`
	AtUserIDList     *[]string        `json:"atUserIDList"`
	MsgDataList      *[]byte          `json:"msgDataList"`
	AttachedInfo     *string          `json:"attachedInfo"`
	Ex               *string          `json:"ex"`
}

type CallbackSingleMsgReadReq struct {
	CallbackCommand `json:"callbackCommand"`
	SendID          string `json:"sendID"`
	ReceiveID       string `json:"receiveID"`
	ContentType     int64  `json:"contentType"`
	EventTime       int64  `json:"eventTime"`
}

type CallbackSingleMsgReadResp struct {
	CommonCallbackResp
}

type CallbackAfterUpdateUserInfoReq struct {
	CallbackCommand `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	UserID          string `json:"userID"`
	Nickname        string `json:"nickName"`
	FaceURL         string `json:"faceURL"`
	Ex              string `json:"ex"`
}
type CallbackAfterUpdateUserInfoResp struct {
	CommonCallbackResp
}

type CallbackBeforeUserRegisterReq struct {
	CallbackCommand `json:"callbackCommand"`
	OperationID     string      `json:"operationID"`
	Secret          string      `json:"secret"`
	Users           []*UserInfo `json:"users"`
	EventTime       int64       `json:"eventTime"`
}

type CallbackBeforeUserRegisterResp struct {
	CommonCallbackResp
	Users []*UserInfo `json:"users"`
}

type CallbackAfterUserRegisterReq struct {
	CallbackCommand `json:"callbackCommand"`
	OperationID     string      `json:"operationID"`
	Secret          string      `json:"secret"`
	Users           []*UserInfo `json:"users"`
	EventTime       int64       `json:"eventTime"`
}

type CallbackAfterUserRegisterResp struct {
	CommonCallbackResp
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID           string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID"`
	Nickname         string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
	FaceURL          string `protobuf:"bytes,3,opt,name=faceURL,proto3" json:"faceURL"`
	Ex               string `protobuf:"bytes,4,opt,name=ex,proto3" json:"ex"`
	CreateTime       int64  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime"`
	AppMangerLevel   int32  `protobuf:"varint,6,opt,name=appMangerLevel,proto3" json:"appMangerLevel"`
	GlobalRecvMsgOpt int32  `protobuf:"varint,7,opt,name=globalRecvMsgOpt,proto3" json:"globalRecvMsgOpt"`
}
