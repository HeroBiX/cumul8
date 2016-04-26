package models

import (
	//Third Party packages
	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		Id            bson.ObjectId `json:"id" bson:"_id"`
		Username      string        `json:"username" bson:"username"`
		UsernameLower string        `json:"usernamelower" bson:"usernamelower"`
		Password      string        `json:"password" bson:"password"`
		Access        string        `json:"access" bson:"access"`
		Filename      []string      `json:"filename" bson"filename"`
	}
)

type (
	// SizeLimit
	SizeLimit struct {
		Limit int64 `json:"limit" bson:"limit"`
	}
)
