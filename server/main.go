package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
		// client.POST("/createProfile", CreateProfile)
		// client.GET("/manageProfile", ManageProfile)
		// client.POST("/manageProfile", ManagedProfile)
		// client.GET("/project/:id", ViewProject)
		// client.POST("/createProject", CreateProject)
	}
	router.Run()
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
