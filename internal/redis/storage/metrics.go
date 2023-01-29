package storage

import (
	"time"

	"redis-timeseries/internal/config"

	redis_timeseries_go "github.com/RedisTimeSeries/redistimeseries-go"
	log "github.com/sirupsen/logrus"
)

var (
	retentionSampleDuration      = config.GetGeneralConfig().RetentionSampleDuration
	retentionAggregationDuration = config.GetGeneralConfig().RetentionAggregationDuration
	aggregationBucketDuration    = config.GetGeneralConfig().AggregationBucketDuration
)

func RedisAddMetric(r *redis_timeseries_go.Client, keyname string, metric float64, labels map[string]string) {
	options := redis_timeseries_go.DefaultCreateOptions
	options.Labels = labels
	options.RetentionMSecs = retentionSampleDuration

	_, err := r.AddAutoTsWithOptions(keyname, metric, options)
	if err != nil {
		log.Error(err)
	}
}

func RedisCreateAVGRule(r *redis_timeseries_go.Client, keyname string, labels map[string]string) {
	labels["aggregator"] = "avg"
	_, haveit := r.Info(keyname + "_avg")
	if haveit != nil {
		options := redis_timeseries_go.DefaultCreateOptions
		options.Labels = labels
		options.RetentionMSecs = retentionAggregationDuration
		err := r.CreateKeyWithOptions(keyname+"_avg", options)
		if err != nil {
			log.Error(err)
			return
		}
		err = r.CreateRule(keyname, redis_timeseries_go.AvgAggregation, uint(aggregationBucketDuration/time.Millisecond), keyname+"_avg")
		if err != nil {
			log.Error(err)
		}
		log.Info("redis TS: rule created ", keyname+"_avg")
	} else {
		err := r.DeleteRule(keyname, keyname+"_avg")
		if err != nil {
			log.Error(err)
			return
		}
		err = r.CreateRule(keyname, redis_timeseries_go.AvgAggregation, uint(aggregationBucketDuration/time.Millisecond), keyname+"_avg")
		if err != nil {
			log.Error(err)
		}
		log.Info("redis TS: rule updated ", keyname+"_avg")
	}
}

func RedisRange(r *redis_timeseries_go.Client, keyname string, fromTimestamp int64, toTimestamp int64) ([]redis_timeseries_go.DataPoint, error) {
	return r.RangeWithOptions(keyname, fromTimestamp, toTimestamp, redis_timeseries_go.DefaultRangeOptions)
}
