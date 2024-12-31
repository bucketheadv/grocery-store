package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"sort"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    *T     `json:"data,omitempty"`
}

func ApiResponseOk[T any](c *gin.Context, response Response[T]) {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}
	ApiResponse(c, response)
}

func ApiResponseError[T any](c *gin.Context, response Response[T]) {
	if response.Code == 0 {
		response.Code = http.StatusInternalServerError
	}
	ApiResponse(c, response)
}

func ApiResponse[T any](c *gin.Context, response Response[T]) {
	c.JSON(response.Code, response)
}

type Person struct {
	Name string
	Age  decimal.Decimal
}

func SortByPeopleAge() {
	people := []Person{
		{"Alice", decimal.NewFromInt(30)},
		{"Bob", decimal.NewFromInt(25)},
		{"Charlie", decimal.NewFromInt(35)},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age.Compare(people[j].Age) < 0
	})

	fmt.Println(people)
}
