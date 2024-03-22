package auth

type AuthService struct {
	ResourceRepo *GftAuthResourceRepo
	RoleRepo     *GftAuthRoleRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		ResourceRepo: NewAuthResourceRepo(),
		RoleRepo:     NewAuthRoleRepo(),
	}
}

func (s *AuthService) SaveResource(id string, input AuthResourceData) (string, error) {
	old, err := s.ResourceRepo.Get(id)
	if err == nil {
		oldMpath := old.GetMpath()
		old.AuthResourceData = input
		_, err = s.ResourceRepo.Save2(old, oldMpath)
		return old.Id, err
	}

	it := NewAuthResource(input)
	_, err = s.ResourceRepo.Save2(it, "")
	return it.Id, err
}
