package models

// CreateLinkModel - model input from user to create a redirect link
type CreateLinkModel struct {
	Link  string `json:"Link"`
	Title string `json:"Title"`
}

// ResponseErrorModel - model for the response i send
type ResponseErrorModel struct {
	Error string `json:"Error"`
}

// StoreDb - stuff to be stored in db
type StoreDb struct {
	Link    string `json:"Link" query:"Link" form:"Link"`
	Title   string `json:"Title" query:"Title" form:"Link"`
	Shorten string `json:"Shorten" query:"Shorten" from:"Shorten"`
}

// SearchDb - an interface for search the db
type SearchDb struct {
	Route string `bson:"shorten"`
}
