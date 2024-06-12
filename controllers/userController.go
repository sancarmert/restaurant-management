package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurant-management/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sancarmert/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10

		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, projectStage,
		})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})

			var allUsers [0]bson.M
			if err = result.All(ctx, &allUsers); err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, allUsers)

		}

		//either pass an error

		//ideally want to return all the users based on the various query parameters

	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("user_id")

		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})

		}
		c.JSON(http.StatusOK, user)

	}
}
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func HashPassword(password string) string {

}
func VerifyPassword(password string, providePassword string) (bool, string) {

}
