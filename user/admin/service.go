package admin

type AdminService struct {
	AdminRepo *GftAdminRepo
}

func NewAdminService() *AdminService {
	return &AdminService{
		AdminRepo: NewAdminRepo(),
	}
}
