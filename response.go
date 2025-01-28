package main

// SuccessResponse allows handlers to specify a custom success status code.
type SuccessResponse struct {
	Status int
	Data   interface{}
}

func NewSuccessResponse(status int, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Status: status,
		Data:   data,
	}
}
