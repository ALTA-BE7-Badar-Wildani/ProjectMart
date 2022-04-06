package user

import (
	web "go-ecommerce/entities/web"
	userRepository "go-ecommerce/repositories/user"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewUserService(repository userRepository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (service UserService) FindAll(limit, page int, filters []map[string]string, sorts []map[string]interface{}) ([]web.UserResponse, error) {

	offset := (page - 1) * limit

	usersRes := []web.UserResponse{}
	users, err := service.userRepo.FindAll(limit, offset, filters, sorts)
	copier.Copy(&usersRes, &users)

	return usersRes, err
}

func (service UserService) Create(userRequest web.UserRequest) (web.UserResponse, error) {
	return web.UserResponse{}, nil
}


func (service UserService) GetPagination(page, limit int) (web.Pagination, error) {
	totalRows, err := service.userRepo.CountAll()
	if err != nil {
		return web.Pagination{}, err
	}
	totalPages :=  totalRows / int64(limit)

	return web.Pagination{
		Page: page,
		Limit: limit,
		TotalPages: int(totalPages),
	}, nil
}