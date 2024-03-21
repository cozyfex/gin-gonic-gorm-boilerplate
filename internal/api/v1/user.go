package v1

import (
	"gin-gonic-gorm-boilerplate/internal/db"
	"gin-gonic-gorm-boilerplate/internal/model"
	"gin-gonic-gorm-boilerplate/internal/repository"
	"gin-gonic-gorm-boilerplate/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	dbManager := c.MustGet("dbManager").(*db.Manager)
	userRepo := repository.NewUserRepository(dbManager)
	userService := service.NewUserService(userRepo)

	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func PostUser(c *gin.Context) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
