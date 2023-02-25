package views

type ResponseTemplate struct {
	Code   int    `json:"code"`   //此处约定：1代表成功，0代表失败
	Msg    string `json:"msg"`    //对请求结果的描述消息，可以为空
	Result any    `json:"result"` //请求成功后给出的结果值
	Err    any    `json:"err"`    //如果请求失败，这里一定要给出错误信息
}
