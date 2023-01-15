package metrics

import (
	"fmt"
	"redis-timeseries/internal/redis"
	"redis-timeseries/internal/redis/storage"
	"reflect"

	log "github.com/sirupsen/logrus"
)

func PublishMetric(keyName string, values any, createRules bool) {
	r := reflect.ValueOf(values).Elem()
	for i := 0; i < r.NumField(); i++ {
		metric := r.Type().Field(i).Name
		metricKeyName := keyName + ":" + metric
		labels := map[string]string{
			"metric": metric,
		}
		if r.Field(i).Kind() == reflect.Slice || r.Field(i).Kind() == reflect.Array {
			s := reflect.ValueOf(r.Field(i).Interface())
			for i := 0; i < s.Len(); i++ {
				keyNameI := fmt.Sprintf("%s:%d", metricKeyName, i)

				value, err := getFloat(s.Index(i).Interface())
				if err != nil {
					log.Error(err)
				} else {
					storage.RedisAddMetric(redis.DB().RedisTS, keyNameI, value, labels)
					if createRules {
						storage.RedisCreateAVGRule(redis.DB().RedisTS, keyNameI, labels)
					}
				}
			}
		} else {
			value, err := getFloat(r.Field(i).Interface())
			if err != nil {
				log.Error(err)
			} else {
				storage.RedisAddMetric(redis.DB().RedisTS, metricKeyName, value, labels)
				if createRules {
					storage.RedisCreateAVGRule(redis.DB().RedisTS, metricKeyName, labels)
				}
			}
		}
	}
}

func getFloat(unk interface{}) (float64, error) {
	var floatType = reflect.TypeOf(float64(0))
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}
