package service

import (
	"encoding/base32"
	"myapp/domain/model"
	"myapp/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ListUsersCommand model.Paginator
type ListUsersResult model.PaginatedUserList

type CreateUserCommand struct {
	Name string
}
type CreateUserResult model.User

type UserService interface {
	ListUsers(c *ListUsersCommand) (*ListUsersResult, error)
	CreateUser(c *CreateUserCommand) (*CreateUserResult, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{
		userRepository: userRepository,
	}
}

func (u userService) ListUsers(c *ListUsersCommand) (*ListUsersResult, error) {

	p := model.Paginator(*c)

	r, err := u.userRepository.List(&p)
	if err != nil {
		return nil, err
	}

	res := ListUsersResult(*r)
	return &res, nil
}

func (u userService) CreateUser(c *CreateUserCommand) (*CreateUserResult, error) {
	now := time.Now()
	newUser := &model.User{
		ID:        generateUserID(),
		Name:      c.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := u.userRepository.Save(newUser)
	if err != nil {
		return nil, err
	}

	res := CreateUserResult(*newUser)
	return &res, nil
}

func generateUserID() string {
	id := uuid.New()
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(id[:])
}
