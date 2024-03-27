package domain

type TxSubscribe struct {
	TxID           string `gorm:"index" json:"tx_id"`
	UserID         string `gorm:"index" json:"user_id"`
	TargetConfirms string `gorm:"index" json:"target_confirms"`
	Status         bool   `gorm:"default:0" json:"status"`
	Model
}
