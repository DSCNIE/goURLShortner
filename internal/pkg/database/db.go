package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/iresharma/REST1/internal/pkg/constants"
	"github.com/iresharma/REST1/internal/pkg/models"
	"github.com/labstack/gommon/random"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conn -  Connecting to the mongoDB
func Conn() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DbString")))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	constants.DbClient = *client
	return true
}

// CreateLink - create the shorten link
func CreateLink(params models.CreateLinkModel) (string, bool) {
	links := constants.DbClient.Database("main").Collection("links")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := links.FindOne(ctx, params)
	_, err := result.DecodeBytes()
	if err != nil {
		store := models.StoreDb{Link: params.Link, Title: params.Title, Shorten: constants.Host + "/short/" + params.Title}
		doc, err := links.InsertOne(ctx, store)
		if err == nil {
			fmt.Printf("%T", doc.InsertedID)
			return constants.Host + "/short/" + params.Title, true
		}
		fmt.Println(err.Error())
		return "", false
	}
	fmt.Println("reached here")
	var short string = random.New().String(7)
	tempStore := models.StoreDb{Link: params.Link, Title: params.Title, Shorten: constants.Host + "/short/" + short}
	_, errnew := links.InsertOne(ctx, tempStore)
	if errnew != nil {
		return "", false
	}
	return constants.Host + "/short/" + short, true
}

// GetLink - getting the real link from the shortened link
func GetLink(seacrhQuery models.SearchDb) (string, bool) {
	links := constants.DbClient.Database("linkShortner").Collection("links")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(seacrhQuery)
	result := links.FindOne(ctx, seacrhQuery)
	model := new(models.StoreDb)
	err := result.Decode(model)
	if err != nil {
		fmt.Println(err.Error())
		return "", false
	}
	return model.Link, true
}
