package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"redis-timeseries/internal/config"
	"redis-timeseries/internal/grpc_timeseries"
	"redis-timeseries/internal/metrics"
	"redis-timeseries/internal/metrics/system"
	"redis-timeseries/internal/router"

	"sync"
	"time"

	"redis-timeseries/internal/redis"
)

var (
	wg       = &sync.WaitGroup{}
	httpPort = config.GetGeneralConfig().HTTPPort
	grpcPort = config.GetGeneralConfig().GRPCPort
)

func main() {
	// Init config
	config.Init_config()

	// Init redis database
	redis.ConnectDB()

	go PublishLoop()
	wg.Add(1)

	grpsTimeSeries := grpc_timeseries.Grpc()
	router := router.InitRouter(grpsTimeSeries.GrpcWebServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}
	go func() {
		// Start Grpc Server
		log.Fatalf("Fail to serve: %v", grpsTimeSeries.GrpcServer.Serve(lis))
	}()
	wg.Add(1)
	go func() {
		// Start HTTP Server
		log.Fatalf("Fail to serve: %v", http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
	}()
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
