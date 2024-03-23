package jc

type GsJcService struct {
	YxsRepo  *GftGsYxsRepo
	ZydmRepo *GftGsZydmRepo
}

func NewGsJcService() *GsJcService {
	return &GsJcService{
		YxsRepo:  NewGsYxsRepo(),
		ZydmRepo: NewGsZydmRepo(),
	}
}

func (s *GsJcService) SaveGsYxs(id string, input GsYxsData) (string, error) {
	old, err := s.YxsRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.GsYxsData = input
		_, err = s.YxsRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	it := NewGsYxs(input)
	_, err = s.YxsRepo.Save2(it, "")
	return it.Id, err
}

func (s *GsJcService) DelGsYxs(id string) (bool, error) {
	return s.YxsRepo.Del(id)
}
