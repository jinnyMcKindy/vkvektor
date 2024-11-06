package main

import (
	fface "facefinder/findface/facerec"
	"gopkg.in/mgo.v2/bson"
	"image"
)

type res struct {
	ID        bson.ObjectId    `json:"_id" bson:"_id,omitempty"`
	URL       string           `json:"url" bson:"url"`
	Rectangle image.Rectangle  `json:"rectangle" bson:"rectangle"`
	Vector    fface.Descriptor `json:"vector" bson:"vector"`
	Img       string           `json:"img,omitempty" bson:"img,omitempty"`
}
