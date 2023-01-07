package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Status  uint16      `json:"status"`
	Data    interface{} `json:"data"`
}

func ReturnJSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type ValidationResponseError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func SuccessResponse(payload interface{}) Response {
	return Response{
		Message: "Request successfully!",
		Status:  200,
		Data:    payload,
	}
}

func InternalServerErrorResponse(payload interface{}) Response {
	return Response{
		Message: "Something went wrong!",
		Status:  500,
		Data:    payload,
	}
}
func BadRequestResponse(payload interface{}, message string) Response {
	return Response{
		Message: message,
		Status:  403,
		Data:    payload,
	}
}
