package helper

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"test-api/model"
)


var hostName  = os.Getenv("MONGO_HOST")
var portNo  = os.Getenv("MONGO_PORT")
var connectionString = "mongodb://"
const dbName = "Netflix"
const colName = "watchList"

var collection *mongo.Collection

//connect to mongoDb

func init(){
	if hostName==""{
		hostName="localhost"
	}
	if portNo==""{
		portNo="27017"
	}
	connectionString = connectionString+hostName+"/"+portNo
	clientOption := options.Client().ApplyURI(connectionString)
	//connect with mongodb
	client,err := mongo.Connect(context.TODO(),clientOption)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)

}

//MONGODB Helper

func InsertOneMovie(movie model.Netflix) *mongo.InsertOneResult{
	inserted,err := collection.InsertOne(context.Background(),movie)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("inserted into db",inserted.InsertedID)
	return inserted
}

//update one record
func UpdateOneMovie(movieId string) int64{
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}
	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	return result.ModifiedCount
}

//delete record

func DeleteOneMovie(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	result,err := collection.DeleteOne(context.Background(),filter,nil)
	if err!=nil{
		log.Fatal(err)
	}
	return result.DeletedCount
}

//delete all movies

func DeleteAllMovies() int64 {
	filter := bson.D{{}}
	result,err := collection.DeleteMany(context.Background(),filter,nil)
	if err!=nil{
		log.Fatal(err)
	}
	return result.DeletedCount
}

//get All movies from database
func GetAllMovies() [] primitive.M{
	filter := bson.D{}
	cursor,err := collection.Find(context.Background(),filter,nil)
	if err!=nil{
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies [] primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

//get oneMovie
func GetOneMovie(movieId string) bson.M{
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	var movie bson.M
	data:= collection.FindOne(context.Background(),filter)
	if data==nil{
		log.Fatal("error")
	}
	data.Decode(&movie)
	return movie
}