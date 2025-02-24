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

type HttpHeader string

func (h HttpHeader) Is(header HttpHeader) bool {
	return h == header
}

func (h HttpHeader) String() string {
	return string(h)
}

var CustomHttpHeader = struct {
	Authorization HttpHeader

	Oss_Bucket HttpHeader
}{
	Authorization: "Authorization",

	Oss_Bucket: "oss_Bucket",
}
