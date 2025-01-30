package internal

type APIResponse struct {
	Status int
	Data   interface{}
}

func NewAPIResponse(status int, data interface{}) *APIResponse {
	return &APIResponse{
		Status: status,
		Data:   data,
	}
}
