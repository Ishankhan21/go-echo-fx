package controllers

import (
	"fmt"

	"github.com/Ishankhan21/go-fiber-fx/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetProfileData(c echo.Context) error {
	// Validate input request etc
	userId := 10 //In real application we extract from request param/body

	userProfile, err := uc.userService.GetUserProfile(userId)
	if err != nil {
		fmt.Errorf("Error %w", err)
	}

	return c.JSON(200, userProfile)
}

func (uc *UserController) ListOrders(c echo.Context) error {
	// Validate input request etc
	userId := 10 //In real application we extract from request param/body

	userOrders, err := uc.userService.GetUsersOrders(userId)
	if err != nil {
		fmt.Errorf("Error %w", err)
	}

	return c.JSON(200, userOrders)
}
