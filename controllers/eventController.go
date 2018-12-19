package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/tobi007/cdServer/repository/dao"
	"github.com/tobi007/cdServer/repository"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/tobi007/cdServer/util"
	"github.com/tobi007/cdServer/models"
	"fmt"
)

func NewEventController(db *gorm.DB) *EventController {
	return &EventController{
		repo: dao.NewMySQLEventRepo(db),
	}
}
type EventController struct{
	repo repository.EventRepo
}

func (evt EventController) NewEvent(c *gin.Context)  {
	var event models.Event

	if err := c.BindJSON(&event); err != nil {
		util.RespondwithJSONAndAbort(c, 400, "bad request")
		return
	}
	fmt.Printf("\nEvent Gotten: %v\n", event)
	_, err := evt.repo.NewEvent(&event)
	if err != nil {
		fmt.Println(err)
		util.RespondwithJSONAndAbort(c, 500, "Error to Creating event")
		return
	}
	util.RespondwithJSON(c, 200, "Events found!",  event)
	return
}

func (evt EventController) GetAllEvent(c *gin.Context) {
	events, err := evt.repo.RetrieveAllEvent()
	if err != nil {
		util.RespondwithJSONAndAbort(c, 500, "Error to retrieve event for user")
		return
	}
	if events != nil {
		c.JSON(200, gin.H{"message": "Events found!", "events": events})
		return
	}
}

func (evt EventController) GetEventByUserId(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		events, err := evt.repo.RetrieveEventByUserId(int64(id))
		if err != nil {
			util.RespondwithJSONAndAbort(c, 500, "Error trying to retrieve events")
			return
		}
		if events != nil {
			util.RespondwithJSON(c, 200, "Events found!",  events)
			return
		}
	}
	util.RespondwithJSONAndAbort(c, 400, "bad request")
	return
}

func (evt EventController) DeleteEvent(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		code, err := evt.repo.DeleteEventByUserId(int64(id))
		if err != nil {
			util.RespondwithJSONAndAbort(c, code, "Error trying to delete events")
			return
		}
		util.RespondwithJSON(c, code, "Events Deleted!", new(models.Event))
		return
	}
	util.RespondwithJSONAndAbort(c, 400, "bad request")
	return
}
