package timeseries

import (
	"redis-timeseries/internal/redis"
	"redis-timeseries/internal/redis/storage"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetTimeSeries(ctx context.Context, in *TimeSerieRequest) (*TimeSeries, error) {
	log.Printf("Receive message from %d to %d", in.FromTimestamp, in.ToTimestamp)
	var dataPoint *DataPoint
	var dataPoints []*DataPoint
	// dataPoint = &DataPoint{Timestamp: in.FromTimestamp, Value: 1111111111111}
	// dataPoints = append(dataPoints, dataPoint)
	// dataPoint = &DataPoint{Timestamp: in.ToTimestamp, Value: 1111111111112}
	// dataPoints = append(dataPoints, dataPoint)

	datapoints, err := storage.RedisRange(redis.DB().RedisTS, in.Key, in.FromTimestamp, in.ToTimestamp)
	if err != nil {
		log.Error(err)
	}

	for i := 0; i < len(datapoints); i++ {
		dataPoint = &DataPoint{Timestamp: datapoints[i].Timestamp, Value: datapoints[i].Value}
		dataPoints = append(dataPoints, dataPoint)
	}

	return &TimeSeries{Key: in.Key, Data: dataPoints}, nil
}
