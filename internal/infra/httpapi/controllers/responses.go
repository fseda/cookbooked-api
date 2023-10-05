package controllers

type ErrResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Data interface{} `json:"data"`
}

var InternalServerErrResponse = &ErrResponse{
	Message: "Internal Server Error",
}
