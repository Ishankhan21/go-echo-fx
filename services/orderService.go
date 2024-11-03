package services

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderService struct {
	redisClient *redis.Client
	mongoClient *mongo.Client
}

func NewOrderService(redisClient *redis.Client, mongoClient *mongo.Client) *orderService {
	return &orderService{redisClient: redisClient, mongoClient: mongoClient}
}

func (s *orderService) GetOrdersByUserId(userId int) ([]map[string]interface{}, error) {
	collection := s.mongoClient.Database("order_management").Collection("orders")
	cur, err := collection.Find(context.Background(), bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}
	var orders []map[string]interface{}
	err = cur.All(context.Background(), &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
