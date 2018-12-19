package models

import "time"

type Event struct {
	ID    		uint 		`json:"id";gorm:"primary_key"`
	Title 		string		`json:"title"`
	Date  		string		`json:"date"`
	Active    	bool   		`json:"active"`
	UserId 		uint		`json:"userId"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"-"`
	DeletedAt 	*time.Time	`json:"-"`
}
