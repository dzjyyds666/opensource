package httpx

type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	data       interface{} `json:"data"`
}
