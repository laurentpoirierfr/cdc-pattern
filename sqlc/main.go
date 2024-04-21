package main

import (
	"bytes"
	"encoding/json"
	"os"
	"sqlc-demo/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

//go:generate sqlc generate

type Query struct {
	ID int32 `uri:"id" binding:"required"`
}

type Response struct {
	Offset int32       `json:"offset"`
	Limit  int32       `json:"limit"`
	Count  int32       `json:"count"`
	Data   interface{} `json:"data"`
}

func main() {
	srv, err := services.NewService()
	OnErr(err)

	router := gin.Default()

	apiCustomers := router.Group("/api/customers")
	{
		apiCustomers.GET("/", func(c *gin.Context) {
			limit := c.DefaultQuery("limit", "100")
			offset := c.DefaultQuery("offset", "0")

			int_limit, err := strconv.ParseInt(limit, 10, 64)
			if err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			int_offset, err := strconv.ParseInt(offset, 10, 64)
			if err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}

			//var limit, offset int32
			customers, err := srv.GetCustomers(int32(int_limit), int32(int_offset))
			if err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			c.JSON(200, Response{
				Data:   customers,
				Offset: int32(int_offset),
				Limit:  int32(int_limit),
				Count:  int32(len(customers)),
			})
		})

		apiCustomers.GET("/:id", func(c *gin.Context) {
			var query Query
			if err := c.ShouldBindUri(&query); err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			customer, err := srv.GetCustomer(query.ID)
			if err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			c.JSON(200, customer)
		})

		apiCustomers.GET("/address/cities/countries/:id", func(c *gin.Context) {
			var query Query
			if err := c.ShouldBindUri(&query); err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			customer, err := srv.GetCustomersByCountry(query.ID)
			if err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			c.JSON(200, customer)
		})

	}

	apiCountries := router.Group("/api/countries")
	{
		apiCountries.GET("/:id", func(c *gin.Context) {
			var query Query
			if err := c.ShouldBindUri(&query); err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			country, err := srv.GetCountry(query.ID)
			if err != nil {
				c.JSON(400, gin.H{"msg": err})
				return
			}
			c.JSON(200, country)
		})
	}

	router.Run("0.0.0.0:" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080

}

func OnErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
