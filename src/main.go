package main

// "github.com/joho/godotenv"
import (
	"context"
	"fmt"
	"go-mongo/modules/api"
	connectdriver "go-mongo/modules/connect-driver"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const secondTimeout = 10

func main() {
	url := "mongodb://adminuser:password123@127.0.0.1:52673"
	mongodatabase := "test-golang"
	collation := "users"
	ctx, _ := context.WithTimeout(context.Background(), secondTimeout*time.Second)
	////
	connect, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	err = connect.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection sucessfull")

	dbx := connectdriver.NewConnection(connect.Database(mongodatabase))
	dbx.ConfigCollation(collation)

	dbxhandler := api.NewTasksHandler(dbx)
	s := api.CreateNewServer()
	s.MountHandlers(dbxhandler)
	fmt.Println("Ready to serve")
	http.ListenAndServe(":3000", s.Router)

}
