package taskshandler

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tasks struct {
	//ID string `bson:"_id,omitempty" json:"id"`
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	//Fecha       time.Time `bson:"fecha" json:"fecha"`
}
