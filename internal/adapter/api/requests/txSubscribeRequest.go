package requests

type CreateTxSubscribeRequest struct {
	TxID string `json:"tx_id" binding:"required"`
	TargetConfirms  int `json:"target_confirms" binding:"required"`
}

type UpdateTxSubscribeRequest struct {
	TargetConfirms  int `json:"target_confirms" binding:"required"`
}