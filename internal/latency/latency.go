package latency

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func TestLatency() {
	rdb := redis.NewClient(&redis.Options{
		//Addr: "redis01.twin24.loc:6379",
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	now := time.Now()
	_, err := rdb.Ping(context.Background()).Result()
	elapsed := time.Since(now)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//fmt.Println("Ping latency: ", float64(elapsed.Microseconds()))
	fmt.Println(float64(elapsed.Microseconds()))

}
