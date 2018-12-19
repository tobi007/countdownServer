package models

import (
	"time"
)

type User struct {
	ID        	uint 		`gorm:"primary_key"`
	Name      	string 		`json:"name"`
	Password  	string 		`json:"-"`
	Active    	bool   		`json:"active,omitempty"`
	CreatedAt 	time.Time 	`json:"-"`
	UpdatedAt 	time.Time 	`json:"-"`
	DeletedAt 	*time.Time	`json:"-"`
}
