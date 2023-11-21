package test

import (
	"bytes"
	"call-back-http/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	url := "http://192.168.62.1:8080/sdkName/callbackBeforeCreateGroupCommand"

	user1 := model.GroupAddMemberInfo{
		UserID:    "user123",
		RoleLevel: 20,
	}
	user2 := model.GroupAddMemberInfo{
		UserID:    "user456",
		RoleLevel: 20,
	}
	user3 := model.GroupAddMemberInfo{
		UserID:    "user789",
		RoleLevel: 20,
	}
	arr := make([]model.GroupAddMemberInfo, 3)
	arr[0] = user1
	arr[1] = user2
	arr[2] = user3
	req := model.CreateGroupReq{
		OperationID:            "123456",
		CallbackCommand:        "callbackBeforeCreateGroupCommand",
		GroupID:                "12345",
		GroupName:              "MyGroup",
		Notification:           "Welcome to MyGroup!",
		Introduction:           "This is a group for discussing example topics.",
		FaceURL:                "http://example.com/path/to/face/image.png",
		OwnerUserID:            "user123",
		CreatorUserID:          "user",
		MemberCount:            10,
		Ex:                     "Extra data",
		Status:                 1,
		CreateTime:             1673048592000,
		GroupType:              1,
		NeedVerification:       1,
		LookMemberInfo:         1,
		ApplyMemberFriend:      0,
		NotificationUpdateTime: 1673048592000,
		NotificationUserID:     "user456",
		GroupAddMemberInfo:     arr,
	}

	byte, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(byte))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respon := &model.CreateGroupResp{}
	err = json.Unmarshal(responseBody, respon)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response:", string(responseBody))
	fmt.Printf("%#v", respon)
	fmt.Println(respon)
}
