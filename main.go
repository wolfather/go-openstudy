package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	firstName string
	lastName string
	age int
}

type People []Person

var people = People{
	Person {
		firstName: "John",
		lastName: "Doe",
		age: 61,
	},
	Person {
		firstName: "Mary",
		lastName: "Hary",
		age: 32,
	},
}


func main() {
	showPeopleData()
	creatServer()
}


func creatServer() {
	server := gin.Default()
	server.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "data": "Hi from GIN" })
	})
	server.Run()
}
