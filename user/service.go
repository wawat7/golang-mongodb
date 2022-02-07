package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-mongodb/exception"
)

type Service interface {
	Create(user User, address Address) User
	FindAll() (users []User)
	FindById(Id string) (user User)
	Update(Id string, request CreateUserRequest) User
	Delete(Id string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(user User, address Address) User {
	user = s.repository.Create(user, address)
	return user
}

func (s *service) FindAll() (users []User) {
	users = s.repository.FindAll()
	return
}

func (s *service) FindById(Id string) (user User) {
	user = s.repository.FindById(Id)
	return user
}

func (s *service) Update(Id string, request CreateUserRequest) User {
	objID, err := primitive.ObjectIDFromHex(Id)
	exception.PanicIfNeeded(err)

	user := User{
		Id:    objID,
		Name:  request.Name,
		Email: request.Email,
		Address: Address{
			Address:  request.Address,
			City:     request.City,
			Province: request.Province,
		},
	}
	user = s.repository.Update(user)

	return user
}

func (s *service) Delete(Id string) {
	s.repository.Delete(Id)

	return
}
