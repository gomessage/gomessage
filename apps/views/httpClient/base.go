package httpClient

import (
    "encoding/json"
    "gomessage/apps/models"
)

// Request 获取用户提交过来的数据
type Request struct {
    *models.Client
    ClientInfo json.RawMessage `json:"client_info"`
}

type Url struct {
    Url string `json:"url"`
}

// ClientInfoDingtalk 存放钉钉客户端ClientInfo中的数据
type ClientInfoDingtalk struct {
    *models.Dingtalk
    RobotUrlList []Url `json:"robot_url"`
}

// ClientInfoFeishu 存放飞书客户端ClientInfo中的数据
type ClientInfoFeishu struct {
    *models.Feishu
    RobotUrlList []Url `json:"robot_url"`
}
