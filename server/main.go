package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var users *mongo.Collection

func init() {
	//start a mongo db session
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	users = client.Database("SoftwareTechnology").Collection("Users")
}

//Client type
type Client struct {
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Gender       string    `json:"gender"`
	DBO          time.Time `json:"dateofbirth"`
	ProfileImage []byte    `json:"profileimage"`
	Description  string    `json:"description"`
	Link         string    `json:"link"`
}

func main() {
	router := gin.Default()

	nonregistered := router.Group("/")
	{
		nonregistered.GET("/", GetHomePage)
		nonregistered.POST("/createProfile", CreateProfile)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/", GetAdminMainPage)
	}

	developer := router.Group("/developer")
	{
		developer.GET("/", GetDevMainPage)
	}

	client := router.Group("/client")
	{
		client.GET("/", GetClientMainPage)
		client.GET("/:email", GetClient)
		// client.GET("/manageProfile", ManageProfile)
		// client.POST("/manageProfile", ManagedProfile)
		// client.GET("/project/:id", ViewProject)
		// client.POST("/createProject", CreateProject)
	}
	router.Run()
}

func GetClient(c *gin.Context) {
	if c.Request.Method != "GET" {
		fmt.Println("Only get here or get out!")
	}
	mail := c.Param("email")

	fmt.Println(mail)

	client := Client{}

	res := users.FindOne(context.Background(), bson.M{"email": mail}).Decode(&client)

	fmt.Println(res)
	fmt.Println(client)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"message": "Email not found"})
	// }
	c.JSON(200, gin.H{
		"message": "user found",
		"data":    client,
	})

}

func CreateProfile(c *gin.Context) {
	if c.Request.Method != "POST" {
		fmt.Println("Only post here no get!")
		return
	}

	email := c.PostForm("email")
	usn := c.PostForm("username")
	pass := c.PostForm("password")
	name := c.PostForm("name")
	surn := c.PostForm("surname")
	gender := c.PostForm("gender")
	// DBO := c.PostForm("DBO")
	// image := c.PostForm("image")
	desc := c.PostForm("description")
	link := c.PostForm("link")

	cl := Client{email, usn, pass, name, surn, gender, time.Now(), []byte{0}, desc, link}

	users.InsertOne(context.Background(), bson.M{}(cl))
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"data":    cl,
	})

}

// GetHomePage : Get the home page for non-registed viewers.
func GetHomePage(c *gin.Context) {
	if c.Request.Method != "GET" {
		fmt.Println("Only get here no give!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Home Page here!",
	})
}

// GetClientMainPage : Get the home page for clients
func GetClientMainPage(c *gin.Context) {
	if c.Request.Method != "GET" {
		fmt.Println("Only get here no give!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Client main Page here!",
	})
}

// GetAdminMainPage : Get the home page for admins
func GetAdminMainPage(c *gin.Context) {
	if c.Request.Method != "GET" {
		fmt.Println("Only get here no give!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Admin main Page here!",
	})
}

// GetDevMainPage : Get the home page for developers
func GetDevMainPage(c *gin.Context) {
	if c.Request.Method != "GET" {
		fmt.Println("Only get here no give!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Dev main Page here!",
	})
}
