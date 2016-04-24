package models

import (
	// Standard Libary packages
	"time"

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
		Online        bool          `json:"online" bson:"online"`
		LastLoggedIn  time.Time     `json:"lastloggedin" bson:"lastloggedin"`
		Access        string        `json:"access" bson:"access"`
		Files         []string      `json:"files" bson"files"`
	}
)
