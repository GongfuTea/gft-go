package cms_handlers

import (
	"github.com/GongfuTea/gft-go/shop/product"
	"github.com/GongfuTea/gft-go/shop/product/commands"
	"github.com/GongfuTea/gft-go/shop/product/queries"
)

type ProductCategoryResolver struct {
	productService *product.ProductService
}

func NewProductCategoryResolver(productService *product.ProductService) *ProductCategoryResolver {
	return &ProductCategoryResolver{
		productService: productService,
	}
}

func (r *ProductCategoryResolver) SaveCmsCategory(cmd commands.SaveProductCategory) (string, error) {
	return r.productService.SaveProductCategory(cmd.Id, cmd.Input)
}

func (r *ProductCategoryResolver) CmsCategories(cmd queries.ProductCategories) ([]*product.GftProductCategory, error) {
	return r.productService.ProductCategoryRepo.All()
}

func (r *ProductCategoryResolver) CmsCategory(q queries.ProductCategory) (*product.GftProductCategory, error) {
	return r.productService.ProductCategoryRepo.Get(q.Id)
}

func (r *ProductCategoryResolver) DelCmsCategory(cmd commands.DelProductCategory) (bool, error) {
	return r.productService.ProductCategoryRepo.Del(cmd.Id)
}
