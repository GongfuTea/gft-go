package pool

import (
	"context"
	"fmt"

	"github.com/GongfuTea/gft-go/x/timex"
	"go.mongodb.org/mongo-driver/bson"
)

type AnalyseXj struct {
}

func (a *AnalyseXj) FindByXxwRq(rq string) []*GftGsXjPoolItem {
	xxwrq := timex.ParseLocalUTC(rq)

	if items, err := GsXjPoolItemRepo.Find(context.Background(), bson.M{"xxwrq": xxwrq}).All(); err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	} else {
		println("count", len(items))
		return items
	}
}
