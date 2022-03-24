package dao

import (
	model "backend-user_crud/model"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type ApiRecordDao struct {
}

func NewApiRecordDao() *ApiRecordDao {
	return &ApiRecordDao{}
}

var (
	redisClient *redis.Client
)

func init() {
	config := GetConfig()

	redisConfig := model.RedisConfig{
		Url:      config.GetString("redis.url"),
		Password: config.GetString("redis.password"),
		Port:     config.GetInt("redis.port"),
		DB:       config.GetInt("redis.db"),
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Url, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Connect redis error:", err)
	}

}

func (apiRecordDao *ApiRecordDao) InitApiRecord() {
	redisClient.Set("Create", 0, 0)
	redisClient.Set("Find", 0, 0)
	redisClient.Set("Update", 0, 0)
	redisClient.Set("Delete", 0, 0)
}

func (apiRecordDao *ApiRecordDao) IncrApiRecord(key string) {
	if times := redisClient.Do("INCR", key); times != nil {
		fmt.Printf("Api %s called times: %d\n", key, times.Val())
	}
}

func (apiRecordDao *ApiRecordDao) GetApiRecord() (rst map[string]interface{}) {
	rst = make(map[string]interface{})

	rst["createTimes"] = redisClient.Do("GET", "Create").Val()
	rst["findTimes"] = redisClient.Do("GET", "Find").Val()
	rst["updateTimes"] = redisClient.Do("GET", "Update").Val()
	rst["deleteTimes"] = redisClient.Do("GET", "Delete").Val()

	fmt.Println(rst)
	return
}
