package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abishek118/momgoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2/bson"
)

const connectionString = "mongodb+srv://mongo_go:go123@cluster0.uxqkb.mongodb.net/?retryWrites=true&w=majority"

const dbName = "Amazingzone"
const colName = "cart"

//MOST IMPORTENT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection instance is ready")

}

// MONGODB helpers - file

//insert 1 record

func insertOneItem(item model.Amazingzone) {
	inserted, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)

}
func updateOneItem(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"item": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

// delete 1 record
func deleteOneItem(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item got delete with delete count:", deleteCount)
}

// delete all record from mongodb

func deleteAllItem() int64 {
	// filter :=
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("no.of items deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount

}

func getAllItems() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies

}

// Actual controller - file

func GetMyAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allItems := getAllItems()
	json.NewEncoder(w).Encode(allItems)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var item model.Amazingzone
	_ = json.NewDecoder(r.Body).Decode(&item)
	insertOneItem(item)
	json.NewEncoder(w).Encode(item)

}

///mark a movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneItem(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneItem(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllItem()
	json.NewEncoder(w).Encode(count)
}
