package services

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	redisClient  *redis.Client
	orderService *orderService
}

func NewUserService(redisConn *redis.Client, orderService *orderService) *UserService {
	return &UserService{redisClient: redisConn, orderService: orderService}
}

func (us *UserService) GetUserProfile(userId int) (any, error) {
	userProfileData, err := us.redisClient.HGetAll(context.Background(), fmt.Sprintf("user:%d", userId)).Result()
	if err != nil {
		return nil, err
	}

	return userProfileData, nil
}

func (us *UserService) GetUsersOrders(userId int) (any, error) {
	orders, err := us.orderService.GetOrdersByUserId(userId)
	if err != nil {
		fmt.Println("Error ", err)
	}

	return orders, nil
}
