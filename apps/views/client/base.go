package client

import (
	"encoding/json"
	"gomessage/apps/models"
	"gomessage/apps/models/clients"
)

type RequestBody struct {
	*models.Client
	ClientInfo json.RawMessage `json:"client_info"`
}

type OneUrl struct {
	Url string `json:"url"`
}

type RequestDataDingtalk struct {
	*clients.Dingtalk
	RobotUrlList []OneUrl `json:"robot_url_list"`
}

type RequestDataFeishu struct {
	*clients.Feishu
	RobotUrlList []OneUrl `json:"robot_url_list"`
}

type RequestDataWechatRobot struct {
	*clients.WechatRobot
	RobotUrlList []OneUrl `json:"robot_url_list"`
}
