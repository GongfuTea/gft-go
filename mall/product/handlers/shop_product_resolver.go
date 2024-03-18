package cms_handlers

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/mall/product"
	"github.com/GongfuTea/gft-go/mall/product/commands"
	"github.com/GongfuTea/gft-go/mall/product/queries"
	"go.mongodb.org/mongo-driver/bson"
)

type ShopProductResolver struct {
	productService *product.ProductService
}

func NewShopProductResolver(productService *product.ProductService) *ShopProductResolver {
	return &ShopProductResolver{
		productService: productService,
	}
}

func (r *ShopProductResolver) SaveShopProduct(cmd commands.SaveShopProduct) (string, error) {
	if cmd.Id != "" {
		_, err := r.productService.ShopProductRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := product.NewShopProduct(cmd.Input)
	_, err := r.productService.ShopProductRepo.Save(item)
	return item.Id, err
}

func (r *ShopProductResolver) ShopProducts(q queries.ShopProducts) (mgo.QueryPageResult[*product.GftShopProduct], error) {
	m := bson.M{}
	if q.Filter.Category != "" {
		m["categoryIds"] = q.Filter.Category
	}
	res, err := r.productService.ShopProductRepo.Find(m).Page(&q.Filter.PagerFilter)

	return res, err

}

func (r *ShopProductResolver) ShopProduct(q queries.ShopProduct) (*product.GftShopProduct, error) {
	return r.productService.ShopProductRepo.Get(q.Id)
}

func (r *ShopProductResolver) DelShopProduct(cmd commands.DelShopProduct) (bool, error) {
	return r.productService.ShopProductRepo.Del(cmd.Id)
}
