package cron

import (
	"github.com/robfig/cron"
)


func TestCronJob() {
	// fmt.Println("Running TestCronJob")
}

func InitCron() {
	// Starting cron job
	c := cron.New()
	c.AddFunc("@every 10s", func() {TestCronJob()})
	c.Start()
}
