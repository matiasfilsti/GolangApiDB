package connectdriver

import (
	"context"
	"fmt"
	"go-mongo/modules/taskshandler"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Db) DbGetAll(ctx context.Context) ([]*taskshandler.Tasks, error) {
	// resp, err := c.link.Collection("users").Find(ctx, map[string]any{})
	fmt.Println(c.coll)

	resp, err := c.link.Collection(c.coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}
	tasks := []*taskshandler.Tasks{}
	if err = resp.All(ctx, &tasks); err != nil {
		log.Fatal(err)
	}
	return tasks, err
}

func (c *Db) GetStats(ctx context.Context) error {

	command := bson.D{{"dbStats", 1}}
	//command := bson.D{{"db.users.find().pretty()", 1}}
	var result *bson.M

	err := c.link.RunCommand(ctx, command).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return nil
}

func (c *Db) DbFindOne(ctx context.Context, nameFilter string) ([]*taskshandler.Tasks, error) {
	filter := bson.M{"name": nameFilter}
	tasks := []*taskshandler.Tasks{}
	//err := c.link.Collection("users").FindOne(ctx, filter).Decode(&tasks)

	// err := c.link.Collection(c.coll).FindOne(ctx, bson.M{}).Decode(&tasks)
	err := c.link.Collection(c.coll).FindOne(ctx, filter).Decode(&tasks)

	if err != nil {

		return nil, err
	}

	return tasks, err
}
