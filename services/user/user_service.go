package user

import (
	"go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"
	userRepository "go-ecommerce/repositories/user"
	"go-ecommerce/utilities"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
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

func (service UserService) Find(id int) (web.UserResponse, error) {
	
	user, err := service.userRepo.Find(id)
	userRes  := web.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}


func (service UserService) Create(userRequest web.UserRequest) (web.AuthResponse, error) {
	
	user := domain.User{}
	copier.Copy(&user, &userRequest)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.AuthResponse{}, web.WebError{ Code: 500, Message: "server error: hashing failed" }
	}
	user.Password = string(hashedPassword)

	user, err = service.userRepo.Store(user)
	if err != nil {
		return web.AuthResponse{}, err
	}

	userRes := web.UserResponse{}
	copier.Copy(&userRes, &user)
	token, err := utilities.CreateToken(user)
	if err != nil {
		return web.AuthResponse{}, err
	}
	authRes := web.AuthResponse{
		Token: token,
		User: userRes,
	}
	return authRes, nil
}


func (service UserService) Update(userRequest web.UserRequest, id int) (web.UserResponse, error) {

	// Find user
	user, err := service.userRepo.Find(id)
	if err != nil {
		return web.UserResponse{}, web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}
	oldHashedPassword := user.Password
	
	// Copy request to found user
	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hash password if request not empty or else
	if userRequest.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			return web.UserResponse{}, web.WebError{ Code: 500, Message: "server error: hashing failed" }
		}
		user.Password = string(hashedPassword)
	} else {
		user.Password = oldHashedPassword
	}

	user, err = service.userRepo.Update(user, id)

	// Convert user domain to user response
	userRes := web.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}


func (service UserService) Delete(id int) error {
	// Find user
	_, err := service.userRepo.Find(id)
	if err != nil {
		return web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}
	
	// Copy request to found user
	err = service.userRepo.Delete(id)
	return err
}