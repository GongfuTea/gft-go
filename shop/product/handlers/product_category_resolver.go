package cms_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/shop/product"
	"github.com/GongfuTea/gft-go/shop/product/commands"
	"github.com/GongfuTea/gft-go/shop/product/queries"
	"github.com/google/uuid"
)

type ProductCategoryResolver struct {
}

func (r *ProductCategoryResolver) SaveCmsCategory(cmd commands.SaveProductCategory) (string, error) {
	if cmd.Id != "" {
		_, err := product.ProductCategoryRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	category := product.GftProductCategory{
		Name:      cmd.Input.Name,
		Note:      cmd.Input.Note,
		SortOrder: cmd.Input.SortOrder,
	}
	category.Id = uuid.NewString()
	category.Pid = cmd.Input.Pid
	category.Code = cmd.Input.Code
	category.CreatedAt = time.Now()
	_, err := product.ProductCategoryRepo.Save(&category)
	return category.Id, err
}

func (r *ProductCategoryResolver) CmsCategories(cmd queries.ProductCategories) ([]*product.GftProductCategory, error) {
	return product.ProductCategoryRepo.All()
}

func (r *ProductCategoryResolver) CmsCategory(q queries.ProductCategory) (*product.GftProductCategory, error) {
	return product.ProductCategoryRepo.Get(q.Id)
}

func (r *ProductCategoryResolver) DelCmsCategory(cmd commands.DelProductCategory) (bool, error) {
	return product.ProductCategoryRepo.Del(cmd.Id)
}
