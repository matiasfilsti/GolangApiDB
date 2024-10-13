package connectdriver

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	link *mongo.Database
	coll string
}

func NewConnection(db *mongo.Database) *Db {
	return &Db{
		link: db,
	}
}

func (f *Db) ConfigCollation(collation string) {
	f.coll = collation
}
