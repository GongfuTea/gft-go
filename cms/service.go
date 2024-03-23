package cms

type CmsService struct {
	BannerCategoryRepo *GftCmsBannerCategoryRepo
	BannerRepo         *GftCmsBannerRepo
	CategoryRepo       *GftCmsCategoryRepo
	NavRepo            *GftCmsNavRepo
	ImageRepo          *GftCmsImageRepo
	PostRepo           *GftCmsPostRepo
}

func NewCmsService() *CmsService {
	return &CmsService{
		BannerCategoryRepo: NewBannerCategoryRepo(),
		BannerRepo:         NewBannerRepo(),
		CategoryRepo:       NewCmsCategoryRepo(),
		NavRepo:            NewCmsNavRepo(),
		ImageRepo:          NewCmsImageRepo(),
		PostRepo:           NewCmsPostRepo(),
	}
}

func (s *CmsService) SaveCmsCategory(id string, input CmsCategoryData) (string, error) {
	old, err := s.CategoryRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.CmsCategoryData = input
		_, err = s.CategoryRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	category := NewCmsCategory(input)
	_, err = s.CategoryRepo.Save2(category, "")
	return category.Id, err
}

func (s *CmsService) DelCmsCategory(id string) (bool, error) {
	return s.CategoryRepo.Del(id)
}

func (s *CmsService) SaveCmsNav(id string, input CmsNavData) (string, error) {
	old, err := s.NavRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.CmsNavData = input
		_, err = s.NavRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	category := NewCmsNav(input)
	_, err = s.NavRepo.Save2(category, "")
	return category.Id, err
}

func (s *CmsService) DelCmsNav(id string) (bool, error) {
	return s.NavRepo.Del(id)
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
