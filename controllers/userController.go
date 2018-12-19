package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tobi007/cdServer/repository/dao"
	"strconv"
	"github.com/tobi007/cdServer/repository"
	"github.com/jinzhu/gorm"
	"github.com/tobi007/cdServer/util"
)

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		repo: dao.NewMySQLUserRepo(db),
	}
}
type UserController struct{
	repo repository.UserRepo
}

func (usr UserController) Retrieve(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		user, err := usr.repo.RetrieveByID(int64(id))
		if err != nil {
			util.RespondwithJSONAndAbort(c, 500, "Error to retrieve data")
			return
		}
		if user != nil {
			util.RespondwithJSON(c, 200, "User founded!",  user)
			return
		}else {
			util.RespondwithJSON(c, 200, "User founded!",  user)
			return
		}
	}
	util.RespondwithJSONAndAbort(c, 400, "bad request")
	return
}

func (u UserController) Create(c *gin.Context)  {

}
