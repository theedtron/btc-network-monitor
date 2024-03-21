package cronjobs

import (
	"btc-network-monitor/internal/adapter/api/rpc"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/database"
	"btc-network-monitor/internal/logger"
	"btc-network-monitor/internal/mailer"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"gorm.io/gorm/clause"
)


func TxNotify (){
	
	//Query txsubscribers
	repo := mysql_repo.TxSubscribeRepository{}
	txModel := repo.ArrayModel()
	db := database.ConnectDB()
	qTx := db.Preload(clause.Associations)
	qTx.Where("status = ?", 0)
	qTx.Find(&txModel)

	client := rpc.Config

	for _, sub := range txModel {
		txId, err := chainhash.NewHashFromStr(sub.TxID)
		if err != nil {
			logger.Error("Error convering chain hash data" + err.Error())
			return
		}
		getTx, err := client.GetTransaction(txId)
		if err != nil {
			logger.Error("Error getting bitcoin tx data" + err.Error())
			return
		}

		confirms, err := strconv.ParseInt(sub.TargetConfirms, 10, 64)
		if err != nil {
			logger.Error("Error convering string to int64" + err.Error())
			return
		}

		if getTx.Confirmations >= confirms {
			userRepo := mysql_repo.UserRepository{}
			userModel := userRepo.Model()
			db := database.ConnectDB()
			qUser := db.Preload(clause.Associations)
			qUser.Where("id = ?", sub.UserID)
			qUser.First(&userModel)

			//send email
			senderEmailData := mailer.EmailData{
				FirstName: "Theed",
				Subject: "BTM Transaction Confirmation",
				MailTo: userModel.Email,
				Confirmations: confirms,
				TxId: sub.TxID,
			}
			mailer.SendEmail(&senderEmailData)
		}
	}
	
}