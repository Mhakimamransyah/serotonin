package users

type UserService struct {
	User_repo Repository
}

func InitUserService(repository Repository) *UserService {
	return &UserService{
		User_repo: repository,
	}
}

type UsersSpec struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Phone    string `validate:"required"`
}

func (service *UserService) RegistersNewUser(users *UsersSpec) error {
	return nil
}

func (service *UserService) Login(username, password string) (*Users, error) {
	return nil, nil
}

func (service *UserService) GetUser(username string) (*Users, error) {
	return nil, nil
}

func (service *UserService) RemoveUser(username string) error {
	return nil
}

func (service *UserService) ModifyUser(user *UsersSpec) error {
	return nil
}
