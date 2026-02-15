package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-fiber/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connection with mongoDb
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb Connected successfully.")
	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}

// Mongodb helpers - file
// insert 1 record

func insertOneMovie(movie model.Netflix) string {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("Something went wrong while adding document", err.Error())
	}

	newId := inserted.InsertedID.(bson.ObjectID).Hex()
	fmt.Println("Inserted 1 movie in db ", newId)
	return newId
}

// update 1
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count: ", result.ModifiedCount)
}

// delete 1 record. multiple
func deleteOneMovie(movieId string) {

	id, _ := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted Count: ", result.DeletedCount)
}

// delete All
func deleteAll() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted Count: ", result.DeletedCount)
	return result.DeletedCount
}

// getAll the movies from database
func getAllMovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie primitive.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// actual controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)

	movieId := params["movieId"]
	updateOneMovie(movieId)
	json.NewEncoder(w).Encode("Movie marked as a watched")
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)

	movieId := params["movieId"]
	deleteOneMovie(movieId)
	json.NewEncoder(w).Encode("Movie deleted")
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAll()
	json.NewEncoder(w).Encode("All the Movie deleted")
}
