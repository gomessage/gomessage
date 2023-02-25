package httpClient

import (
    "encoding/json"
    "gomessage/apps/models"
)

// RequestBody 获取用户提交过来的数据
type RequestBody struct {
    *models.Client
    ClientInfo json.RawMessage `json:"client_info"`
}

type OneUrl struct {
    Url string `json:"url"`
}

// ClientInfoDingtalk 存放钉钉客户端ClientInfo中的数据
type ClientInfoDingtalk struct {
    *models.Dingtalk
    RobotUrlList []OneUrl `json:"robot_url_list"`
}

// ClientInfoFeishu 存放飞书客户端ClientInfo中的数据
type ClientInfoFeishu struct {
    *models.Feishu
    RobotUrlList []OneUrl `json:"robot_url_list"`
}
