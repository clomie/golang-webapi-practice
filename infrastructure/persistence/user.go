package persistence

import (
	"myapp/domain/model"
	"myapp/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(entity *model.User) error {
	tx := r.db.Create(entity)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *userRepository) List(p *model.Paginator) (*model.PaginatedUserList, error) {

	var users []model.User
	r.db.Limit(p.Limit).Offset(p.Offset).Find(&users)

	var count int64
	r.db.Model(&model.User{}).Count(&count)

	total := int(count)

	res := model.PaginatedUserList{
		Items:  users,
		Total:  total,
		Offset: p.Offset,
		Limit:  p.Limit,
	}

	return &res, nil
}

func (r *userRepository) Get(ID string) (*model.User, error) {
	panic("not implemented")
}

func (r *userRepository) Update(entity *model.User) error {
	panic("not implemented")
}

func (r *userRepository) Delete(ID string) (*model.User, error) {
	panic("not implemented")
}
