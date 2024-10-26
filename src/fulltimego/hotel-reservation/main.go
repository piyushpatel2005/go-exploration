package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/piyushpatel2005/hotel-reservation/api"
	"github.com/piyushpatel2005/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const uri = "mongodb://root:password@localhost:27017"
const dbname = "hotel-reservation"
const userCollection = "users"

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "server listen address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	// verify the connection to mongodb works
	//fmt.Println(client)
	//ctx := context.Background()
	//userCol := client.Database(dbname).Collection(userCollection)
	//user := types.User{
	//	FirstName: "John",
	//	LastName:  "Smith",
	//}
	//_, err = userCol.InsertOne(ctx, user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var james types.User
	//if err := userCol.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(james)

	app := fiber.New(config)

	apiv1 := app.Group("/api/v1")
	app.Get("/foo", handleFoo)

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Hello there",
	})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "john.smith",
	})
}
