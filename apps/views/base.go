package views

type ResponseTemplate struct {
	Msg string `json:"msg"`
	Err any    `json:"err"`
}
