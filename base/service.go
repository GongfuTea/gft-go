package base

type BaseService struct {
	DictCategoryRepo *GftDictCategoryRepo
	DictItemRepo     *GftDictItemRepo
}

func NewBaseService() *BaseService {
	return &BaseService{
		DictCategoryRepo: NewDictCategoryRepo(),
		DictItemRepo:     NewDictItemRepo(),
	}
}

func (s *BaseService) SaveDictCategory(id string, input DictCategoryData) (string, error) {
	old, err := s.DictCategoryRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.DictCategoryData = input
		_, err = s.DictCategoryRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	category := NewDictCategory(input)
	_, err = s.DictCategoryRepo.Save2(category, "")
	return category.Id, err
}

func (s *BaseService) DelDictCategory(id string) (bool, error) {
	return s.DictCategoryRepo.Del(id)
}
