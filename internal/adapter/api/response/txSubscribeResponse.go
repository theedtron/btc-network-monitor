package response

import (
	"btc-network-monitor/internal/helper"
)

type TxSubscribeDataResponse map[string]interface{}
type TxSubscribeDataArrayResponse []map[string]interface{}

func NewTxSubscribeResponse(data interface{}, err error) Response {
	jsonResp, _ := helper.ToJson(data)

	return Response{
		Message: "successful",
		Data:    TxSubscribeDataResponse(jsonResp),
	}
}

func NewTxSubscribeArrayResponse(data interface{}, err error) Response {
	jsonResp, _ := helper.ToArrayJson(data)

	return Response{
		Message: "successful",
		Data:    TxSubscribeDataArrayResponse(jsonResp),
	}
}