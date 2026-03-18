package main

// Entry point for the HTTP API -> starts a gin server
import (
	"fmt"

	"github.com/Giovanni-Hernandez10/ZillaRCE/internal/sandbox"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dockerClient, err := sandbox.CreateDockerClient("python:alpine", "python")
	if err != nil {
		fmt.Print("Error setting up docker client")
	}
	if dockerClient != nil {
		fmt.Printf("Created docker client")
	}
	fmt.Printf("The docker client was successfully created")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message: ": "pong"})
	})
	r.Run()
}
