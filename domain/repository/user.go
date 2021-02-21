package repository

import "myapp/domain/model"

type UserRepository interface {
	Save(entity *model.User) error
	List(p *model.Paginator) (*model.PaginatedUserList, error)
	Get(ID string) (*model.User, error)
	Update(entity *model.User) error
	Delete(ID string) (*model.User, error)
}
