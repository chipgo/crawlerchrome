package main

import (
	"context"
	"log"
	"time"

	"crawlerchrome/config"
	"crawlerchrome/utils"

	"github.com/chromedp/chromedp"
	"github.com/robfig/cron/v3"
)

func main() {

	cfg := config.NewConfig()

	opts := []cron.Option{cron.WithLocation(utils.HCMLocationTime)}
	cronJob := cron.New(opts...)

	if _, err := cronJob.AddFunc(cfg.CronSchedule.Crawler, crawlGoDev); err != nil {
		log.Println("Failed to add crawlGoDev job")
	}

	cronJob.Start()
	time.Sleep(15 * time.Second)
	cronJob.Stop()
}

func crawlGoDev() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// true : for host machine debug to show chrome tabs
		// false : for dev or prod
		chromedp.Flag("headless", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var example string
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
