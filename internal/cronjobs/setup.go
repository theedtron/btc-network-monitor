package cronjobs

import (
	"btc-network-monitor/internal/logger"
	"btc-network-monitor/internal/mailer"
	"github.com/robfig/cron"
)

func InitCronJob(stop <-chan struct{}) {
	c := cron.New()
	mailr := mailer.NotificationConfig{}
	err := c.AddFunc("1 * * * *", func() {
		logger.Info("Tx notify running...")
		TxNotify(mailr)
	})
	if err != nil {
		logger.Error("Failed to add cron job:" + err.Error())
		return
	}
	c.Start()
	logger.Info("Cron job started")

	<-stop

	c.Stop()

}
