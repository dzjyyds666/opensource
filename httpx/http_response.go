package httpx

import (
	"github.com/labstack/echo"
	"time"
)

type HttpResponse struct {
	Data       interface{} `json:"data"`
	StatusCode StatusCode  `json:"status_code"`
	Ts         int64       `json:"ts"`
	Url        string      `json:"url"`
}

func JsonResponse(ctx echo.Context, statusCode StatusCode, data interface{}) error {
	return ctx.JSON(statusCode.Int(), HttpResponse{
		StatusCode: statusCode,
		Data:       data,
		Ts:         time.Now().Unix(),
		Url:        ctx.Request().URL.Path,
	})
}
