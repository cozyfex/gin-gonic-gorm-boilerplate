package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/model"
	"gin-gonic-gorm-boilerplate/internal/repository"
	"gin-gonic-gorm-boilerplate/internal/service"
)

func ListUser(c *gin.Context, db *database.Manager) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	users, err := userService.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(c *gin.Context, db *database.Manager) {
	userRepo := repository.NewUserRepository(db)
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

func User(c *gin.Context, db *database.Manager) {
	email := c.Param("email")
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	user, err := userService.User(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
