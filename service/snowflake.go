//SNOWFLAKE算法实现，调用"github.com/gitstliu/go-id-worker"算法
package service

import (
	"github.com/gitstliu/go-id-worker"
	"github.com/spf13/viper"
)

var worker *idworker.IdWorker

//初始化
func init() {
	workerId := viper.GetInt64("snowflake.worker-id")
	dataCenterId := viper.GetInt64("snowflake.data-center-id")
	worker = &idworker.IdWorker{}
	if err := worker.InitIdWorker(workerId, dataCenterId); err != nil {
		panic(err.Error())
	}
}

//获取主键
func GetId() (int64, error) {
	return worker.NextId()
}

//获取多个主键
func GetIds(count int) ([]int64, error) {
	var ids []int64
	for i := 0; i < count; i++ {
		if id, err := GetId(); err == nil {
			ids = append(ids, id)
		} else {
			return nil, err
		}
	}
	return ids, nil
}
