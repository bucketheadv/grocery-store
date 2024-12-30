package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func ApiResponseOK(c *gin.Context, h gin.H) {
	c.JSON(http.StatusOK, h)
}

type Person struct {
	Name string
	Age  int
}

func SortByPeopleAge() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
