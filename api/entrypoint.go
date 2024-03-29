package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	app             *gin.Engine
	mongoClient     *mongo.Client
	mongoClientOnce sync.Once
	taskCollection  *mongo.Collection
)

func connectToMongoDB() {
	mongoClientOnce.Do(func() {
		var err error
		mongoCtx := context.Background()
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)
		tlsConfig := &tls.Config{}
		opts.SetTLSConfig(tlsConfig)

		//clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI")) // Use environment variable
		fmt.Print("MONGO URL: ", os.Getenv("MONGODB_URI"))
		mongoClient, err = mongo.Connect(mongoCtx, opts)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
		if err := mongoClient.Ping(mongoCtx, nil); err != nil {
			log.Printf("Error pinging MongoDB: %v", err)
		}
		taskCollection = mongoClient.Database("ManageTasks").Collection("Tasks")
	})
}

// Create Endpoints
func GetTasks(r *gin.RouterGroup) {
	r.GET("/get", GetTasksIdHandler)
}
func CreateTask(r *gin.RouterGroup) {
	r.POST("/create", CreateTaskHandler)
}
func UpdateTask(r *gin.RouterGroup) {
	r.PUT("/update/:id", UpdateTaskHandler)
}
func DeleteTask(r *gin.RouterGroup) {
	r.DELETE("/delete/:id", DeleteTaskHandler)
}
func GetTask(r *gin.RouterGroup) {
	r.GET("/get/:id", GetTaskByIdHandler)
}

func init() {

	connectToMongoDB()
	gin.SetMode(gin.ReleaseMode)
	app = gin.New()
	app.Use(cors.Default())
	r := app.Group("/api")
	GetTasks(r)
	CreateTask(r)
	UpdateTask(r)
	DeleteTask(r)
	GetTask(r)

	// defer func() {
	// 	mongoClient.Disconnect(context.Background())
	// }()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
