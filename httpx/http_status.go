package httpx

type StatusCode int

func (sc StatusCode) Is(statusCode StatusCode) bool {
	return sc == statusCode
}

func (sc StatusCode) Int() int16 {
	return int16(sc)
}

//const (
//	HttpOK            StatusCode = 200
//	HttpBadRequest    StatusCode = 400
//	HttpNotFound      StatusCode = 404
//	HttpInternalError StatusCode = 500
//	HttpUnauthorized  StatusCode = 401
//	HttpForbidden     StatusCode = 403
//)

var HttpStatusCode = struct {
	HttpOK            StatusCode
	HttpBadRequest    StatusCode
	HttpNotFound      StatusCode
	HttpInternalError StatusCode
	HttpUnauthorized  StatusCode
	HttpForbidden     StatusCode
}{
	HttpOK:            200,
	HttpBadRequest:    400,
	HttpNotFound:      404,
	HttpInternalError: 500,
	HttpUnauthorized:  401,
	HttpForbidden:     403,
}

type HttpHeader string

func (h HttpHeader) Is(header HttpHeader) bool {
	return h == header
}

func (h HttpHeader) String() string {
	return string(h)
}

var CustomHttpHeader = struct {
	Authorization HttpHeader

	Oss_Bucket    HttpHeader
	Oss_Directory HttpHeader
	Oss_RegionId  HttpHeader
}{
	Authorization: "Authorization",

	Oss_Bucket:    "Oss_Bucket",
	Oss_Directory: "Oss_Directory",
	Oss_RegionId:  "Oss_RegionId",
}
