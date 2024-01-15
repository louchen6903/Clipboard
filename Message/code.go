package codes

// ReturnMsg 返回信息
type ReturnMsg struct {
	Code int    `json:"code"` // 错误代码
	Msg  string `json:"msg"`  // 错误信息
}

var (
	/****************** 正常状态：0 ******************/

	DefaultSuccess = &ReturnMsg{Code: 0, Msg: "Success"}

	/****************** 服务错误：1000 ******************/
	ErrService = &ReturnMsg{Code: 1000, Msg: "服务错误:%v"}
)
