package httpx

type StatusCode int

func (sc StatusCode) Is(statusCode StatusCode) bool {
	return sc == statusCode
}

const (
	HttpOK            StatusCode = 200
	HttpBadRequest    StatusCode = 400
	HttpNotFound      StatusCode = 404
	HttpInternalError StatusCode = 500
	HttpUnauthorized  StatusCode = 401
	HttpForbidden     StatusCode = 403
)
