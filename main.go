package main

import (
	_ "fmt"
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
	creatServer()
}


func creatServer() {
	router := gin.Default()
	router.GET("/home", homeUseCase)

	router.Run("3000")
}

func homeUseCase(ctx *gin.Context) {
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, gin.H{ "data": homeDataProvider(name)})
}


func homeDataProvider(name string) string {
	return "Hi there, "+ name+"!"
}