// routes.go
package api

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)




type Tasks struct {
	Id          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	IsCompleted bool     `json:"isCompleted" bson:"isCompleted"`
	CreatedAt   string   `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string   `json:"updatedAt" bson:"updatedAt"`
	CreatedBy   string   `json:"author" bson:"author"`
	Comments    []string `json:"comments" bson:"comments"` // Comments field
	Shared      bool     `json:"shared" bson:"shared"`     // Share field
}

// CreateTaskHandler handles the creation of a new task.
func CreateTaskHandler(ctx *gin.Context) {

	var task Tasks
	if err := ctx.BindJSON(&task); err != nil {
		log.Fatal(err.Error())
		ctx.JSON(400, gin.H{"message": "Error in Parsing Request Body"})
		return
	}

	// Generate a random ID for the task
	task.Id = generateRandomID()

	_, err := taskCollection.InsertOne(ctx, task)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{"message": "Error in Creating Task"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Task Created Successfully", "data": task})
}

// UpdateTaskHandler handles the update of an existing task.
func UpdateTaskHandler(ctx *gin.Context) {

	taskID := ctx.Param("id")
	var updatedTask Tasks
	updatedTask.Id = taskID
	if err := ctx.BindJSON(&updatedTask); err != nil {
		log.Fatal(err.Error())
		ctx.JSON(400, gin.H{"message": "Error in Parsing Request Body"})
		return
	}

	// Additional functionality: Update updatedAt field
	updatedTask.UpdatedAt = time.Now().String()

	filter := bson.M{"id": taskID}
	update := bson.M{"$set": updatedTask}

	result, err := taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{"message": "Error in Updating Task"})
		return
	}

	if result.ModifiedCount == 0 {
		ctx.JSON(404, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Task Updated Successfully"})
}

// DeleteTaskHandler handles the deletion of an existing task.
func DeleteTaskHandler(ctx *gin.Context) {
	taskID := ctx.Param("id")
	filter := bson.M{"id": taskID}
	result, err := taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{"message": "Error in Deleting Task"})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(404, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Task Deleted Successfully"})
}

// GetTaskByIdHandler retrieves a task by its ID.
func GetTaskByIdHandler(ctx *gin.Context) {
	taskID := ctx.Param("id")
	var task Tasks
	filter := bson.M{"id": taskID}

	err := taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(404, gin.H{"message": "Task not found"})
		}
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{"message": "Error in Fetching Task"})
	}

	ctx.JSON(200, gin.H{"message": "Success", "data": task})
}

// GetTaskByIdHandler retrieves a task by its ID.
func GetTasksIdHandler(ctx *gin.Context) {
	log.Println("Request for Get Tasks Arrived")
	var response []Tasks
	cur, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(200, gin.H{"message": "Error in Fetching the Tasks"})
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var task Tasks
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal(err.Error())
			ctx.JSON(200, gin.H{"message": "Error in Fetching the Tasks"})
		}
		response = append(response, task)
	}

	// Return the Response
	ctx.JSON(200, struct {
		Message string  `json:"message"`
		Data    []Tasks `json:"data"`
	}{
		Message: "Success",
		Data:    response,
	})

}

func generateRandomID() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000) + 1
	randomID := fmt.Sprintf("%d", randomNumber)
	return randomID
}
