package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/tobi007/cdServer/repository"
	"github.com/tobi007/cdServer/models"
	"github.com/pkg/errors"
	"fmt"
)

func NewMySQLEventRepo(Conn *gorm.DB) repository.EventRepo {
	return &mysqlEventRepo{
		Conn: Conn,
	}
}
type mysqlEventRepo struct {
	Conn *gorm.DB
}

func (eventRepo *mysqlEventRepo) RetrieveAllEvent() (*[]models.Event, error) {
	var events []models.Event
	if err := eventRepo.Conn.Find(&events).Error; err != nil {
		return nil, nil
	}
	return &events, nil
}

func (eventRepo *mysqlEventRepo) RetrieveEventByUserId(id int64) (*[]models.Event, error) {
	var events []models.Event
	if err := eventRepo.Conn.Where("user_id=?", id).Find(&events).Error; err != nil {
		return nil, nil
	}
	return &events, nil
}

func (eventRepo *mysqlEventRepo) NewEvent(evt *models.Event) (int64, error) {
	evt.Active = true
	res := eventRepo.Conn.Save(evt)
	return res.RowsAffected, res.Error
}

func (eventRepo *mysqlEventRepo) DeleteEventByUserId(id int64) (int, error) {
	var event models.Event
	eventRepo.Conn.First(&event, id)

	if event.ID != 0 {
		eventRepo.Conn.Delete(&event)
		return 200, nil
	} else {
		return 404, errors.New("Event not found")
	}
}
