package cronjobs

import (
	"btc-network-monitor/internal/adapter/api/rpc"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/logger"
	"btc-network-monitor/internal/mailer"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"strconv"
)

func TxNotify(config mailer.NotificationConfig) {
	logger.Info("Starting Tx Notify Check")

	statusArray, _ := mysql_repo.NewTxSubscribeRepository().GetFalseStatus()

	subscribes := statusArray.([]domain.TxSubscribe)

	client := rpc.Config

	for _, sub := range subscribes {
		txHash, err := chainhash.NewHashFromStr(sub.TxID)
		if err != nil {
			logger.Error("Error converting chain hash data" + err.Error())
			return
		}
		getTx, err := client.GetTransaction(txHash)
		if err != nil {
			logger.Error("Error getting bitcoin tx data" + err.Error())
			return
		}

		confirms, err := strconv.ParseInt(sub.TargetConfirms, 10, 64)
		if err != nil {
			logger.Error("Error converting string to int64" + err.Error())
			return
		}

		if getTx.Confirmations >= confirms {
			logger.Info("Preparing to send Email. Target met")

			userRepo := mysql_repo.UserRepository{}
			user, _ := userRepo.Find(sub.UserID)
			userModel, ok := user.(domain.User)
			if !ok {
				logger.Error("failed to type assert user")
				return
			}

			subscriberRepo := mysql_repo.TxSubscribeRepository{}
			subData, err := subscriberRepo.Find(sub.ID)
			if err != nil {
				return
			}

			subscribe, ok := subData.(domain.TxSubscribe)
			if !ok {
				logger.Error("failed to type assert subscribe")
				return
			}

			originalData := subData
			subscribe.Status = true

			_, err = subscriberRepo.Update(sub.ID, subscribe)
			if err != nil {
				logger.Error("failed to update data: " + err.Error())
				return
			}

			//send email
			senderEmailData := mailer.EmailData{
				FirstName:     userModel.Firstname,
				Subject:       "BTM Transaction Confirmation",
				MailTo:        userModel.Email,
				Confirmations: confirms,
				TxId:          sub.TxID,
			}
			err = config.SendEmail(&senderEmailData)
			if err != nil {
				_, rollbackErr := subscriberRepo.Update(sub.ID, originalData)
				if rollbackErr != nil {
					logger.Error("failed to update data: " + err.Error())
				}
				return
			}
		}
	}

}
