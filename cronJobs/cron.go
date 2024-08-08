package cronJobs

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCronJob() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { // Every minute
		fmt.Println("Cron job running:", time.Now())
		// Place your scheduled task here

	})
	c.Start()
}
