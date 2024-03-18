package cms

type CmsService struct {
	BannerCategoryRepo *GftCmsBannerCategoryRepo
	BannerRepo         *GftCmsBannerRepo
}

func NewCmsService() *CmsService {
	return &CmsService{
		BannerCategoryRepo: NewBannerCategoryRepo(),
		BannerRepo:         NewBannerRepo(),
	}
}

func (s *CmsService) SaveBannerCategory(id string, input CmsBannerCategoryData) (string, error) {
	old, err := s.BannerCategoryRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.CmsBannerCategoryData = input
		_, err = s.BannerCategoryRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	category := NewBannerCategory(input)
	_, err = s.BannerCategoryRepo.Save2(category, "")
	return category.Id, err
}

func (s *CmsService) DelBannerCategory(id string) (bool, error) {
	return s.BannerCategoryRepo.Del(id)
}
