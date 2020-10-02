package module

import (
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	ClassCol *mongo.Collection
)

func init() {

	//client, err := mongo.NewClient(options.Client().
	//	ApplyURI())

}