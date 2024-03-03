package cms_handlers

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/shop/product"
	"github.com/GongfuTea/gft-go/shop/product/commands"
	"github.com/GongfuTea/gft-go/shop/product/queries"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type ShopProductResolver struct {
}

func (r *ShopProductResolver) SaveShopProduct(cmd commands.SaveShopProduct) (string, error) {
	if cmd.Id != "" {
		_, err := product.ShopProductRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := product.GftShopProduct{
		// Note:      cmd.Input.Note,
		// SortOrder: cmd.Input.SortOrder,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	_, err := product.ShopProductRepo.Save(&item)
	return item.Id, err
}

func (r *ShopProductResolver) ShopProducts(q queries.ShopProducts) (mgo.QueryPageResult[*product.GftShopProduct], error) {
	m := bson.M{}
	if q.Filter.Category != nil {
		if *q.Filter.Category == "" {
			m["categoryIds"] = []string{}
		} else {
			m["categoryIds"] = *q.Filter.Category
		}
	}
	res, err := product.ShopProductRepo.Find(context.Background(), m).Page(&q.Filter.PagerFilter)

	return res, err

}

func (r *ShopProductResolver) ShopProduct(q queries.ShopProduct) (*product.GftShopProduct, error) {
	return product.ShopProductRepo.Get(q.Id)
}

func (r *ShopProductResolver) DelShopProduct(cmd commands.DelShopProduct) (bool, error) {
	return product.ShopProductRepo.Del(cmd.Id)
}
