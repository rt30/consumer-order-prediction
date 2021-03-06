package main

import (
	"github.com/consumer-order-prediction/pkg/csv"
	orderspb "github.com/consumer-order-prediction/pkg/proto/orders"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy",
	})
}

//Return most popular restaurant with the help of grpc
func  GetPoplarRestaurant(c *gin.Context) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	defer conn.Close()

	client := orderspb.NewOrderServiceClient(conn)

	req := &orderspb.GetPopularRestaurantRequest{
	}

	res, err := client.GetPopularRestaurant(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	c.JSON(200, gin.H{
		"Most Popular Restaurant": res.Name,
	})
}

func  GetPopularVegCuisine(c *gin.Context) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	defer conn.Close()

	client := orderspb.NewOrderServiceClient(conn)

	req := &orderspb.GetPopularVegCuisineRequest{
	}

	res, err := client.GetPopularVegCuisine(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	c.JSON(200, gin.H{
		"Most Popular Veg Cuisine is": res.Name,
	})
}


//Return a specific order on the basis of customer id with the help of gRPC
func GetSpecificOrdersByQuery(c *gin.Context) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	defer conn.Close()

	client := orderspb.NewOrderServiceClient(conn)

	customerid := c.Query("CustomerID")

	req := &orderspb.GetSpecificOrderRequest{
		OrderId:customerid,
	}

	res, err := client.GetSpecificOrder(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	if err!=nil{
		c.JSON(200, gin.H{
			"message":"customer not found",
		})
	}else {
		ginRes := &csv.Order{
			CustomerID:res.GetOrder().CustomerId,
			CustomerName:res.GetOrder().CustomerName,
			RestsurantName:res.GetOrder().RestsurantName,
			VegCuisine:res.GetOrder().VegCuisine,
			NonVegCuisine:res.GetOrder().NonvegCuisine,
			State:res.GetOrder().State,
		}
		c.JSON(200,ginRes)
	}
}

func main(){
	router := gin.Default()

	api:= router.Group("/api",gin.BasicAuth(gin.Accounts{
		"team1": "team1",
	}))
	// http://localhost:5656/api/
	api.GET("/",  HomePage)
	api.GET("/popularrestaurant", GetPoplarRestaurant)
	api.GET("/popularcuisine", GetPopularVegCuisine)
	api.GET("/orders", GetSpecificOrdersByQuery)

	router.Run("localhost:5656")
}