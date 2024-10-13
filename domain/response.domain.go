package domain

type SuccessResponse struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"message"`
}

type AppResult struct {
	Data interface{}
	Message string
	Err error
	StatusCode int
}
