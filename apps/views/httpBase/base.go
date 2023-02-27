package httpBase

type ResponseTemplate struct {
	Code   int    `json:"code"`   //此处约定：1代表成功，0代表失败
	Msg    string `json:"msg"`    //对请求结果的描述消息，可以为空
	Result any    `json:"result"` //请求成功后给出的结果值
	Error  any    `json:"error"`  //如果请求失败，这里一定要给出错误信息
	Help   string `json:"help"`
}

func ResponseSuccessful(msg string, result any) ResponseTemplate {
	return ResponseTemplate{
		Code:   1,
		Msg:    msg,
		Help:   "帮助信息请查看：https://github.com/gomessage/gomessage/blob/master/wiki/response.md ",
		Result: result,
	}
}

func ResponseFailure(msg string, error any) ResponseTemplate {
	return ResponseTemplate{
		Code:  0,
		Msg:   msg,
		Help:  "帮助信息请查看：https://github.com/gomessage/gomessage/blob/master/wiki/response.md ",
		Error: error,
	}
}
