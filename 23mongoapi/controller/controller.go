package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/drunkonbytes/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://asinha:arjuns@cluster0.1rjvuhg.mongodb.net/?retryWrites=true&w=majority"
const dbNa me = "netflix"
const colName = "watchlist"

// reference to the original collection
var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	//collection instance
	fmt.Println("Collection reference is ready")
}

//MongoDB helpers

//insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserter, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with ID: ", inserter.InsertedID)
}

//update 1 record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}
