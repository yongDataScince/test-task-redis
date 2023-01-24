package main

import (
	"context"
	"fmt"
	"os"
	"redis-test/pkg/service"
	"redis-test/pkg/setupdb"
	"redis-test/pkg/utils"
)

func main() {
	redisUrl := os.Getenv("REDIS_URL");
	fmt.Println("Redis on: ", redisUrl);
	// Init main context
	ctx := context.Background()
	currDate := utils.GetCurrDate()
	requestUrl := fmt.Sprintf("https://api.punkapi.com/v2/beers?brewed_before=%s&abv_gt=6", currDate)

	// Init redis DB
	rdb, err := setupdb.NewSetup(&ctx)
	if err != nil {
		errMsg := utils.FormatErrorMsg("Setup redis DB", err)
		panic(errMsg)
	}

	srv := service.NewService(requestUrl, &ctx, rdb)

	if err := srv.MakeRequest(); err != nil {
		errMsg := utils.FormatErrorMsg("Make request", err)
		panic(errMsg)
	}
	fmt.Println("Set values to redis successful")
}
