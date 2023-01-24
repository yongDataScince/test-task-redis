package service_test

import (
	"context"
	"fmt"
	"redis-test/pkg/service"
	"redis-test/pkg/setupdb"
	"redis-test/pkg/utils"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	ctx := context.Background()
	currDate := utils.GetCurrDate()
	requestUrl := fmt.Sprintf("https://api.punkapi.com/v2/beers?brewed_before=%s&abv_gt=6", currDate)

	rdb, err := setupdb.NewSetup(&ctx)
	if err != nil {
		panic(err)
	}

	srv := service.NewService(requestUrl, &ctx, rdb);
	if err := srv.MakeRequest(); err != nil {
		panic(err)
	}

	value, err := rdb.Get(ctx, "Rabiator").Result();
	if err != nil {
		panic(err)
	}

	if value != "Columbus, Hersbrucker, Extra Pale, Wheat, Crystal" {
		t.Error("Value \"Rabiator\" not set in redis");
	}
}