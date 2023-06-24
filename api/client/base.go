package client

import (
	"encoding/json"
	"gomessage/models"
	clients2 "gomessage/models/clients"
)

type RequestBody struct {
	*models.Client
	ClientInfo json.RawMessage `json:"client_info"`
}

type OneUrl struct {
	Url string `json:"url"`
}

type RequestDataDingtalk struct {
	*clients2.Dingtalk
	RobotUrlList []OneUrl `json:"robot_url_list"`
}

type RequestDataFeishu struct {
	*clients2.Feishu
	RobotUrlList []OneUrl `json:"robot_url_list"`
}

type RequestDataWechatRobot struct {
	*clients2.WechatRobot
	RobotUrlList []OneUrl `json:"robot_url_list"`
}
