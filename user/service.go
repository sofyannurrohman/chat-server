package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface{
	RegisterUser(input RegisterUserInput)(User,error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func(s *service) RegisterUser(input RegisterUserInput) (User,error){
	user:= User{}
	user.ID = uuid.NewString()
	user.Name = input.Name
	user.Email = input.Email
	passwordHash,err:= bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash) 
	user.Role = "user"
	newUser,err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser,nil

}

