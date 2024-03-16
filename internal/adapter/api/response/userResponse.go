package response

import (
	"btc-network-monitor/internal/helper"
)

type UserDataResponse map[string]interface{}
type UserDataArrayResponse []map[string]interface{}

func NewUserResponse(data interface{}, err error) Response {
	jsonResp, _ := helper.ToJson(data)

	return Response{
		Message: "successful",
		Data:    UserDataResponse(jsonResp),
	}
}

func NewUserArrayResponse(data interface{}, err error) Response {
	jsonResp, _ := helper.ToArrayJson(data)

	return Response{
		Message: "successful",
		Data:    UserDataArrayResponse(jsonResp),
	}
}
