package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	server := NewServer(Config{port: ":8080"})
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func helloRouteHandler(c *gin.Context) (interface{}, error) {
	return gin.H{
		"message": "Hello World",
	}, nil
}
