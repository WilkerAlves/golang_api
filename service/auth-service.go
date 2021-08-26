package service

import (
	"github.com/wilker/golang_api/dto"
	"github.com/wilker/golang_api/entity"
	"github.com/wilker/golang_api/repository"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential (email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userReporitory repository.UserRepository
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	panic("implement me")
}

func (service *authService) CreateUser(user dto.UserCreateDTO) entity.User {
	panic("implement me")
}

func (service *authService) FindByEmail(email string) entity.User {
	panic("implement me")
}

func (service *authService) IsDuplicateEmail(email string) bool {
	panic("implement me")
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userReporitory: userRep,
	}
}
