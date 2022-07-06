package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

//Define WRapper aroun redis Client
type StorageService struct {
	redisclient *redis.Client
}

//TOP notif for storeservice and redis content
var (
	storeservice = &StorageService{}
	ctx =context.Background()
)
//int storeservice
func InitializeStore() *StorageService{
	redisClient := redis.Newclient(&redis.Options{
		Addr: "Localhost:6379",
		Password: "",
		DB:			0,
	})
	pong,err := redisClient.ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("ERR init redis: %v",err))
	}
	fmt.Printf("\nRedis started:pong Message ={%s}",pong)
	storeService.redisClient = redisClient
	return storeservice


	//maping between orginal and Generated URL

	func SaveUrlMapping(shortUrl string, originalUrl string, userId string){



}
//retrieve longurl from shorted

func RetrieveInitialUrl(shortUrl string) string {


}

func SaveUrlMapping(shortUrl string ,orginalUrl string, userId string){
	err:= storeService.redisClent.set(ctx,shortUrl,orginalUrl ,CacheDuration).ERR()
	if err != nil {
		panic(fmt.Sprintf("Failed recovery URL | ERR: %v - shortURL :%s\n",err,shortUrl))
	}
	return result
}

}

