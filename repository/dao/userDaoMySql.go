package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/tobi007/cdServer/models"
	"github.com/tobi007/cdServer/repository"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewMySQLUserRepo(Conn *gorm.DB) repository.UserRepo {
	return &mysqlUserRepo{
		Conn: Conn,
	}
}

type mysqlUserRepo struct {
	Conn *gorm.DB
}

func (userRepo *mysqlUserRepo)Create( u *models.User) (int64, error) {
	return 0, nil
}

func (userRepo *mysqlUserRepo) Update(u *models.User) (*models.User, error) {
	return new(models.User), nil
}

func (userRepo *mysqlUserRepo) RetrieveByID (id int64) (*models.User, error)  {
	var user models.User
	if err := userRepo.Conn.Where("id=?", id).First(&user).Error; err != nil {
		return nil, nil
	}
	return &user, nil
}

func (userRepo *mysqlUserRepo)ConnectionActive() error  {
	return userRepo.Conn.DB().Ping()
}

