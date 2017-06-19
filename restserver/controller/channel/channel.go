package channel

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mkfsn/shyuukan-program/restserver/model"
)

type controller struct {
	db *gorm.DB
}

func newController(db *gorm.DB) *controller {
	return &controller{
		db: db,
	}
}

func (ctrl *controller) findID(id string) ([]model.Channel, error) {
	var channels []model.Channel

	// Get all matched records
	ctrl.db.Where("id = ?", id).Find(&channels)

	return channels, nil
}

func (ctrl *controller) list(c *gin.Context) {
	var channels []model.Channel
	ctrl.db.Find(&channels)
	c.JSON(http.StatusOK, channels)
}

func (ctrl *controller) create(c *gin.Context) {
	now := time.Now()
	channel := model.Channel{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Owner:       "",
		CreatedAt:   now,
		ModifiedAt:  now,
	}
	log.Printf("[Channel][Create] %v\n", channel)

	if channel.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data for creating channel"})
		return
	}

	if ok := ctrl.db.NewRecord(channel); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops"})
		return
	}

	res := ctrl.db.Create(&channel)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Channel has been created successfully"})
}

func (ctrl *controller) update(c *gin.Context) {
	var id int
	var err error
	var channel model.Channel

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid channel id"})
		return
	}

	if res := ctrl.db.First(&channel, id); res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Channel id not found"})
		return
	}

	channel.Name = c.PostForm("name")
	channel.Description = c.PostForm("description")
	channel.ModifiedAt = time.Now()

	if res := ctrl.db.Save(&channel); res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Channel has been updated successfully"})
}

func (ctrl *controller) delete(c *gin.Context) {
	var id int
	var err error
	var channel model.Channel

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid channel id"})
		return
	}

	res := ctrl.db.First(&channel, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Channel id not found"})
		return
	}

	ctrl.db.Delete(&channel)
	c.JSON(http.StatusOK, gin.H{"message": "Channel has been deleted successfully"})
}

func Routes(relativePath string, route *gin.Engine, db *gorm.DB) {
	controller := newController(db)
	channels := route.Group(relativePath)
	{
		channels.GET("/", controller.list)
		channels.POST("/", controller.create)
		channels.PUT("/:id", controller.update)
		channels.DELETE("/:id", controller.delete)
	}
}
