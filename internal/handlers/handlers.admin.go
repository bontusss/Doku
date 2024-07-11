package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAdmin(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}