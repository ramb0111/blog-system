package article

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SUCCESS = "SUCCESS"
)

func ErrResponse(code int, err error) Response {
	return Response{
		Status:  code,
		Message: err.Error(),
	}
}

func SuccessResponse(code int, data interface{}) Response {
	return Response{
		Status:  code,
		Message: SUCCESS,
		Data:    data,
	}
}
