package cron

import (
	"log"

	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/scrapper"
	"github.com/robfig/cron"
)

// InitCron start our cronjob
func InitCron() {
	// Starting cron job
	c := cron.New()
	c.AddFunc("@every 20m", func() {
		if err := scrapper.UpdateShares(); err != nil {
			log.Fatalln(err)
		}
	})
	c.Start()
}
