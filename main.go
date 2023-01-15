package main

import (
	"redis-timeseries/internal/config"
	"redis-timeseries/internal/metrics"
	"redis-timeseries/internal/metrics/system"

	"sync"
	"time"

	"redis-timeseries/internal/redis"
)

var wg = &sync.WaitGroup{}

func main() {
	// Init config
	config.Init_config()

	// Init redis database
	redis.ConnectDB()

	go PublishLoop()

	wg.Add(1)
	wg.Wait()

}

func PublishLoop() {
	ticker := time.NewTicker(config.GetGeneralConfig().PublishRate)
	quit := make(chan struct{})
	createRules := true
	for {
		select {
		case <-ticker.C:
			go func(createRules bool) {
				//Add new metrics to publish here:
				systemMetrics := system.GetSystemStatus()
				metrics.PublishMetric("system", &systemMetrics, createRules)
				//
			}(createRules)
			createRules = false
		case <-quit:
			ticker.Stop()

			defer wg.Done()
			return
		}
	}
}
