package cron

import (
	"job-funnel/tasks"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()
	c.AddFunc("0 8,12,16 * * *", func() {
		tasks.Weworkremotely_comRss()
	})
	c.Start()
}
