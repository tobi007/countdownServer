package db

import (
	"fmt"
	"github.com/tobi007/cdServer/config"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tobi007/cdServer/models"
)

var db *gorm.DB
var err error


// ConnectSQL ...
func Init() {
	c := config.GetConfig()
	dbSource := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		c.Get("DB_HOST"),
		c.Get("DB_PASS"),
		c.Get("DB_NAME"),
	)
	db, err = gorm.Open("mysql", dbSource)
	if err != nil {
		log.Println("Failed to connect to database ", err)
		panic(err)
	}



	fmt.Println("Connected to Database succefully")
	if c.Get("DROP") != true {
		dropCreatedTable()
		db.AutoMigrate(models.Event{}, &models.User{})
		if c.Get("SEED") == true {
			seedTable()
		}
	}

}

func seedTable()  {
	log.Println("Seeding the user and event table")
	//db.Model(&models.Event{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE") // Foreign key need to define manually
	db.Create(&models.User{
		Name: "Kayode Emmanuel",
		Password: "password",
		Active: true,
	})
	db.Create(&models.Event{
		Title: "Demo a new app",
		Date: "2018-12-29T13:45:18.000Z",
		Active: true,
		UserId: 1,
	})
	db.Create(&models.Event{
		Title: "Go on holiday",
		Date: "2019-05-18T08:09:24.000Z",
		Active: true,
		UserId: 1,
	})
}

func dropCreatedTable()  {
	log.Println("Dropping all tables")
	db.DropTable(models.Event{}, &models.User{})
}

func GetDB() *gorm.DB {
	return db
}

// ConnectMgo ....
func ConnectMgo(host, port, uname, pass string) error {
	return nil
}

