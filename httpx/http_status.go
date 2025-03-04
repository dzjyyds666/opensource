package httpx

type StatusCode int

func (sc StatusCode) Is(statusCode StatusCode) bool {
	return sc == statusCode
}

func (sc StatusCode) Int() int {
	return int(sc)
}

var HttpStatusCode = struct {
	HttpOK            StatusCode
	HttpBadRequest    StatusCode
	HttpParamsError   StatusCode
	HttpNotFound      StatusCode
	HttpInternalError StatusCode
	HttpUnauthorized  StatusCode
	HttpForbidden     StatusCode
}{
	HttpOK:            200,
	HttpBadRequest:    400,
	HttpNotFound:      404,
	HttpParamsError:   422,
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
