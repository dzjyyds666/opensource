package httpx

import "github.com/labstack/echo"

type HttpResponse struct {
	StatusCode StatusCode  `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func JsonResponse(ctx echo.Context, statusCode StatusCode, msg *string, data ...interface{}) error {
	return ctx.JSON(statusCode.Int(), HttpResponse{
		statusCode,
		*msg,
		data,
	})
}
