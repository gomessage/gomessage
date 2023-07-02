package utils

type ResponseTemplate struct {
	Code   int    `json:"code"`   //此处约定：1代表成功，0代表失败
	Msg    string `json:"msg"`    //对请求结果的描述消息，可以为空
	Result any    `json:"result"` //如果请求成功，这里给出成功的结果
	Error  any    `json:"error"`  //如果请求失败，这里一定要给出错误的信息
	Help   string `json:"help"`   //显示接口文档地址，便于别人排错
}

func ResponseSuccessful(msg string, result any) ResponseTemplate {
	return ResponseTemplate{
		Code:   1,
		Msg:    msg,
		Result: result, //响应成功要把result附上
		Help:   "帮助信息请查看：https://github.com/gomessage/gomessage/blob/master/wiki/response.md ",
	}
}

func ResponseFailure(msg string, error any) ResponseTemplate {
	return ResponseTemplate{
		Code:  0,
		Msg:   msg,
		Error: error, //响应失败要把error附上
		Help:  "帮助信息请查看：https://github.com/gomessage/gomessage/blob/master/wiki/response.md ",
	}
}
