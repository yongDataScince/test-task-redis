package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"redis-test/pkg/types"
	"redis-test/pkg/utils"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	mx  *sync.Mutex
	rdb *redis.Client
	ctx *context.Context

	RequestUrl string
}

func NewService(reqUrl string, ctx *context.Context, rdb *redis.Client) *Service {
	return &Service{
		mx:         &sync.Mutex{},
		rdb:        rdb,
		ctx: 				ctx,
		RequestUrl: reqUrl,
	}
}

func (s *Service) MakeRequest() error {
	var data []types.Beer
	response, err := http.Get(s.RequestUrl)

	if err != nil {
		errMsg := utils.FormatErrorMsg("ERROR IN MAKE REQUEST", err);
		return fmt.Errorf(errMsg)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errMsg := utils.FormatErrorMsg("ERROR IN READ BODY", err);
		return fmt.Errorf(errMsg)
	}

	strungify := string(body)
	if len(strungify) == 0 {
		return fmt.Errorf("NO DATA IN RESPONSE")
	}

	if err := json.Unmarshal([]byte(strungify), &data); err != nil {
		errMsg := utils.FormatErrorMsg("ERROR IN MARSHAL DATA", err);
		return fmt.Errorf(errMsg)
	}

	for _, drink := range data {
		var ingNames []string

		for _, hop := range drink.Ingredients.Hops {
			ingNames = append(ingNames, hop.Name)
		}
		for _, m := range drink.Ingredients.Malt {
			ingNames = append(ingNames, m.Name)
		}

		values := strings.Join(ingNames, ", ")
		fmt.Printf("set \"%s\" with value \"[%s]\"\n\n", drink.Name, values);

		if err := s.rdb.Set(*s.ctx, drink.Name, values, 0).Err(); err != nil {
			errMsg := utils.FormatErrorMsg("ERROR IN SET DATA REDIS", err);
			return fmt.Errorf(errMsg)
		}
	}

	return nil
}
