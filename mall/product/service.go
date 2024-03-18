package product

type ProductService struct {
	ProductCategoryRepo *GftProductCategoryRepo
	ShopProductRepo     *GftShopProductRepo
}

func NewProductService() *ProductService {
	return &ProductService{
		ProductCategoryRepo: NewProductCategoryRepo(),
		ShopProductRepo:     NewShopProductRepo(),
	}
}

func (s *ProductService) SaveProductCategory(id string, input ProductCategoryData) (string, error) {
	old, err := s.ProductCategoryRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.ProductCategoryData = input
		_, err = s.ProductCategoryRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	category := NewProductCategory(input)
	_, err = s.ProductCategoryRepo.Save2(category, "")
	return category.Id, err
}

func (s *ProductService) DelProductCategory(id string) (bool, error) {
	return s.ProductCategoryRepo.Del(id)
}
