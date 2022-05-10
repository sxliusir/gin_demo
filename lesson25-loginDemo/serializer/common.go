package serializer

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg interface{} `json:"msg"`
	Error string `json:"error"`
}