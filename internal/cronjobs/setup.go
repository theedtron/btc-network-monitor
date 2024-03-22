package cronjobs

import(
	"time"
	"github.com/robfig/cron"
)


func InitCronJob() {
	//Add cron job
	c := cron.New()
	// Define the Cron job schedule
    c.AddFunc("15 * * * *", func() {
        TxNotify()
    })
	// Start the Cron job scheduler
    c.Start()
	// Wait for the Cron job to run
    time.Sleep(5 * time.Minute)

    // Stop the Cron job scheduler
    c.Stop()
}