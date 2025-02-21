package httpx

type HttpResponse struct {
	StatusCode StatusCode  `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}
