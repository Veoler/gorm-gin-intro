package gormginintro

import (
	"gorm.io/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetStudentByIDHand(c *gin.Context, db *gorm.DB) {
    id, errID := strconv.Atoi(c.Param("id")) // Получаем ID из URL
    if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID неправильный"})
	}
	
	var req Student

    if err := db.First(&req, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Студент не найден"})
        return
    }

    c.JSON(http.StatusOK, req)
}


func CreateStudentHand(c *gin.Context, db *gorm.DB) {
    var req struct {
		Name		string	`json:"name" binding:"required"`
		Age			int		`json:"age" binding:"required"`
	}

    // 1. Валидация JSON
    if errBody := c.ShouldBindJSON(&req); errBody != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Тело запроса не верно"})
        return
    }

    // 2. Создание записи в БД
    student := Student{Name: req.Name, Age: req.Age}
    if err := db.Create(&student).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать студента"})
        return
    }

    // 3. Ответ
    c.JSON(http.StatusCreated, student)
}

func GetStudentsHand(c *gin.Context, db *gorm.DB) {
    var reqs []Student
    
    // Находим всех студентов
    if err := db.Find(&reqs).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reqs)
}

func DeleteStudentHand(c *gin.Context, db *gorm.DB) {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	
	result := db.Delete(&Student{}, id)

	if result.Error !=  nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
   		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Студент не найден"})
   		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func UpdateStudentHand(c *gin.Context, db *gorm.DB) {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	var input struct {Name string `json:"name" binding:"required"`}
	if errBody := c.ShouldBindJSON(&input); errBody != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Тело запроса неправильное"})
        return
    }

	var student Student
	if err := db.First(&student, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Студент не найден"})
        return
    }

	if errUpd := db.Model(&student).Update("Name", input.Name).Error; errUpd != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка обновления"})
		return
	}

	c.JSON(http.StatusOK, student)
}



