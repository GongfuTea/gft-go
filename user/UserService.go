package user

type UserService struct {
	UserRepo *GftUserRepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: NewUserRepo(),
	}
}
