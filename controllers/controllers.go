package controllers

import (
	"context"
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	guuid "github.com/google/uuid"
)

var collection *mongo.Collection

func EmployeesCollection(c *mongo.Database) {
	collection = c.Collection("employees")
}

type Employee struct {
	ID string			`json:"id"`
    FirstName string	`json:"firstName"`
    LastName  string	`json:"lastName"`
	Occupation string	`json:"occupation"`
	SalaryGrade string	`json:"salaryGrade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// creates employee from the data passed into the request
func CreateEmployee(c *gin.Context) {
	var employee Employee
	c.BindJSON(&employee)
	firstName := employee.FirstName
	lastName := employee.LastName
	occupation := employee.Occupation
	salaryGrade := employee.SalaryGrade
	id := guuid.New().String()

	newEmployee := Employee{
		ID: id,
		FirstName:		firstName,
		LastName:		lastName,
		Occupation:		occupation,
		SalaryGrade: 	salaryGrade,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newEmployee)
	if err != nil {
		log.Printf("Error while inserting new employee into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, newEmployee)
	return
}

// gets employee based on the employeeId
func GetEmployee(c *gin.Context) {
	id := c.Param("id")

	employee := Employee{}
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&employee)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Employee not found",
		})
		return
	}

	c.JSON(http.StatusOK, employee)
	return
}
