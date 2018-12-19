package repository

import "github.com/tobi007/cdServer/models"

type UserRepo interface {
	Create( u *models.User) (int64, error)
	Update( u *models.User) (*models.User, error)
	RetrieveByID (id int64) (*models.User, error)
}

type EventRepo interface {
	RetrieveAllEvent() (*[]models.Event, error)
	RetrieveEventByUserId(id int64) (*[]models.Event, error)
	NewEvent(evt *models.Event) (int64, error)
	DeleteEventByUserId(id int64) (int, error)
}
