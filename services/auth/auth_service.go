package auth

import (
	userRepository "go-ecommerce/repositories/user"
	"go-ecommerce/utilities"

	web "go-ecommerce/entities/web"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewAuthService(userRepo userRepository.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}


func (service AuthService) Login(authReq web.AuthRequest) (web.AuthResponse, error) {
	
	// Get user by username
	user, err := service.userRepo.FindBy("username", authReq.Username)
	if err != nil {
		return web.AuthResponse{}, web.WebError{ Code: 401, Message: "Invalid credential" }
	}
	
	userRes := web.UserResponse{}
	copier.Copy(&userRes, &user)

	// Verify password
	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))
	if match != nil {
		return web.AuthResponse{}, web.WebError{ Code: 401, Message: "Invalid password" }
	}

	// Create token
	token, err := utilities.CreateToken(user)
	if err != nil {
		return web.AuthResponse{}, web.WebError{ Code: 500, Message: "Error create token" }
	}
	return web.AuthResponse{
		Token: token,
		User: userRes,
	}, nil
}