package constants

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DbClient - mongo db client
var DbClient mongo.Client

// Host - host name
var Host string = "localhost:5000"
