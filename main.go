package main

import (
	"github.com/Veoler/gorm-gin-intro/gormginintro"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"log"
)

func main() {
    dsn := "host=localhost user=postgres password=4545 dbname=gorm-gin-intro port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Ошибка подключения: %v", err)
    }

    // Автомиграция
    if err := db.AutoMigrate(&gormginintro.Student{}, &gormginintro.Group{}); err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong"})
	})
	r.POST("/students", func(c *gin.Context){
		gormginintro.CreateStudentHand(c, db)
	})
	r.GET("/students/:id", func(c *gin.Context){
		gormginintro.GetStudentByIDHand(c, db)
	})
	r.GET("/students/", func(c *gin.Context){
		gormginintro.GetStudentsHand(c, db)
	})
	r.DELETE("students/:id", func (c *gin.Context){
		gormginintro.DeleteStudentHand(c, db)
	})
	r.PATCH("/students/:id", func (c *gin.Context){
		gormginintro.UpdateStudentHand(c, db)
	})
	r.POST("/groups", func(c *gin.Context){
		gormginintro.CreateGroupHand(c, db)
	})
	r.GET("/groups", func(c *gin.Context){
		gormginintro.GetGroupsHand(c, db)
	})
	r.GET("/groups/:id", func(c *gin.Context){
		gormginintro.GetGroupByIDHand(c, db)
	})
	r.DELETE("/groups/:id", func(c *gin.Context){
		gormginintro.DeleteGroupHand(c, db)
	})
	r.PATCH("/groups/:id", func (c *gin.Context){
		gormginintro.UpdateGroupHand(c, db)
	})
	
	r.Run()
}
