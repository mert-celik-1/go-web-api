package main

import (
	"fmt"
	"go-web-api/src/config"
	"go-web-api/src/infra/cache"
	"time"
)

func main() {
	fmt.Println("!... Hello World ...!")

	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()

	if err != nil {
		fmt.Print(err)
	}

	cli := cache.GetRedis()

	cache.Set(cli, "as2", "test2", time.Hour)

	as, _ := cache.Get[string](cli, "as2")

	fmt.Println(as)

}
