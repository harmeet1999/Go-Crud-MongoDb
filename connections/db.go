package connections

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The init function will run before our main function to establish a connection to MongoDB. If it cannot connect it will fail and the program will exit.
func init() {
	if err := ConnectDb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

// A global variable that will hold a reference to the MongoDB client
var DB *mongo.Database

func ConnectDb() error {
	mongoErr := godotenv.Load()

	if mongoErr != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	DB = client.Database("Products")

	return err
}
