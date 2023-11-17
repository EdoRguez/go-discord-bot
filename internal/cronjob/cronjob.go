package cronjob

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

type CronJob struct {
	CronJob *cron.Cron
}

func (c *CronJob) StartCronJob(spec string, action func()) error {
	if len(c.CronJob.Entries()) > 0 {
		return fmt.Errorf("there is already %v CronJob(s) running", len(c.CronJob.Entries()))
	}

	c.CronJob.AddFunc(spec, action)
	c.CronJob.Start()
	return nil
}

func (c *CronJob) StopCronJobs() {
	if c.CronJob != nil {
		c.CronJob.Stop()
		c.CronJob = NewCronJob()
	}
}

func NewCronJob() *cron.Cron {
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		fmt.Println(err)
	}

	return cron.NewWithLocation(loc)
}
