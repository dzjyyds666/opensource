package httpx

import (
	"github.com/labstack/echo"
)

type HttpResponse struct {
	StatusCode StatusCode  `json:"status_code"`
	Data       interface{} `json:"data"`
}

func JsonResponse(ctx echo.Context, statusCode StatusCode, data interface{}) error {
	return ctx.JSON(statusCode.Int(), HttpResponse{
		statusCode,
		data,
	})
}
