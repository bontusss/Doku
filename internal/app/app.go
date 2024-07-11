package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type App struct {
	Name     string
	Version  string
	Router   *gin.Engine
	DB       *mongo.Client
	Handlers map[string]gin.HandlerFunc
}

func NewApp(name, version, dbUri string) *App {
	// Initialize Mongo client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB successfully connected...")

	return &App{
		Name: name,
		Version: version,
		DB: client,
		Router: gin.Default(),
	}
}
