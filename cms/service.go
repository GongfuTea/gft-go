package cms

type CmsService struct {
	BannerCategoryRepo *GftCmsBannerCategoryRepo
}

func NewCmsService() *CmsService {
	return &CmsService{
		BannerCategoryRepo: NewBannerCategoryRepo(),
	}
}

func (s *CmsService) SaveBannerCategory(id string, input CmsBannerCategoryData) (string, error) {
	old, err := s.BannerCategoryRepo.Get(id)
	if err == nil {
		old.CmsBannerCategoryData = input
		_, err = s.BannerCategoryRepo.Save2(old, old.GetMpath())
		return old.Id, err
	}

	category := NewBannerCategory(input)
	_, err = s.BannerCategoryRepo.Save2(category, "")
	return category.Id, err
}

func (s *CmsService) DelBannerCategory(id string) (bool, error) {
	return s.BannerCategoryRepo.Del(id)
}
