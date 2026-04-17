package gormginintro

import (
	"net/http"
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func CreateGroupHand (c *gin.Context, db *gorm.DB) {
	var req struct {
		Name	string	`json:"name" binding:"required"`
	}

	if errBody := c.ShouldBindJSON(&req); errBody != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Тело запроса неверно"})
		return
	}

	group := Group{Name: req.Name}
	if err := db.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}

	c.JSON(http.StatusCreated, group)
}

func GetGroupsHand(c *gin.Context, db *gorm.DB) {
	var req []Group

	if err := db.Find(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}

	c.JSON(http.StatusOK, req)
}

func GetGroupByIDHand(c *gin.Context, db *gorm.DB) {
	id, errID := strconv.Atoi(c.Param("id")) 
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID неправильный"})
		return
	}

	var group Group
	if err := db.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Группа не найдена"})
		return
	}

	c.JSON(http.StatusOK, group)
}

func DeleteGroupHand(c *gin.Context, db *gorm.DB) {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID неправильный"})
		return
	}

	var group Group
	if err := db.Delete(&group, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}

	if db.Delete(&group, id).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Группа не найдена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func UpdateGroupHand(c *gin.Context, db *gorm.DB) {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID неправильный"})
		return
	}

	var input struct {
		Name	string	`json:"name" binding:"required"`	
	}
	if errBody := c.ShouldBindJSON(&input); errBody != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Тело запроса не верно"})
		return
	}

	group := Group{Name: input.Name}
	if err := db.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Группа не найдена"})
		return
	} 

	if errUpd := db.Model(&group).Update("Name", input.Name).Error; errUpd != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка обновления"})
		return
	}

	c.JSON(http.StatusOK, group)

}
