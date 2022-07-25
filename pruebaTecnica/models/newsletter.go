package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Newsletter ...
type Newsletter struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string             `bson:"Name" json:"Name"`
	SentEmails int32              `bson:"SentEmails" json:"SentEmails"`
	Recipients []string           `bson:"Recipients" json:"Recipients"`
}

//Email ..
type Email struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Subject string             `bson:"Name" json:"Name"`
	to      []string           `bson:"SentEmails" json:"SentEmails"`
	name    string             `bson:"Recipients" json:"Recipients"`
}

//CreateNewsletter ...
type CreateNewsletter struct {
	Name string `form:"name" json:"name" binding:"required,max=100"`
}

//UpdateNewsletter ...
type UpdateNewsletter struct {
	IDNewsletter   string `bson:"_id" json:"_id"`
	EmailRecipient string `bson:"EmailRecipient" json:"EmailRecipient"`
}

//SendNewsletter ...
type SendNewsletter struct {
	Content string `bson:"Content" json:"Content"`
	Subject string `bson:"Subject" json:"Subject"`
}

//File ...
type StoreFile struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Path     string             `bson:"Path" json:"Path"`
	FileName string             `bson:"FileName" json:"FileName"`
}
