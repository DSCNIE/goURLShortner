package constants

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// DbClient - mongo db client
var DbClient mongo.Client

// Host - host name
var Host string = os.Getenv("host")

// var Host string = "http://127.0.0.1:5000"
