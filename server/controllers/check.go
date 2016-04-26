package controllers

import (
	// Standard Library packages
	"fmt"
	"regexp"
	"strings"

	// Thirt Party packages
	"github.com/HeroBiX/cumul8/server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var b bool

func CheckCreatingUser(un, pass, lun string, uc UserController) bool {
	mgSession := uc.session
	fmt.Printf(": mgSession %T \n", mgSession)

	// See if username looks good
	switch {
	case EnoughCharacters(un, 2) == false:
		StatusHTML = "I do think you need more characters in your username... just a feeling"
		b = false

	case NoFunkyCharacters(un) == false:
		StatusHTML = "The unicorns sense funky characters in your UserName"
		b = false

	case ExistingUserName(un, lun, mgSession) == false:
		StatusHTML = "Sorry sir, the Unicorns have assigned that username to someone else... much cooler than you"
		b = false

	default:

		// See if password looks good
		switch {
		case EnoughCharacters(pass, 2) == false:
			StatusHTML = "The hamsters need more characters in your password.. at least 3"
			b = false

		case NoFunkyCharacters(pass) == false:
			StatusHTML = "The hamster doesn't like funky characters in your password"
			b = false

		default:
			b = true
		}
	}
	return b

}

func EnoughCharacters(a string, z int) bool {
	// check if username and password has enough characters
	if len(a) <= z {
		b = false
	} else {
		b = true
	}
	return b
}

var validPath = regexp.MustCompile("[a-zA-Z0-9]")

// check if user are using symbols
func NoFunkyCharacters(a string) bool {
	// loop through a to see if it has any symbols
	for c, _ := range a {
		if m := validPath.FindStringSubmatch(a[c : c+1]); m == nil {
			b = false
			break
		}
		b = true
	}
	return b
}

// find out if the username already exist
func ExistingUserName(un, lun string, mg *mgo.Session) bool {
	u := models.User{}

	// find if the username exist
	mg.DB("file-server").C("users").Find(bson.M{"username": un}).One(&u)

	if un == u.Username {
		b = false
	} else {

		// find is lower case username exist
		mg.DB("file-server").C("users").Find(bson.M{"usernamelower": lun}).One(&u)

		if lun == u.UsernameLower {
			b = false
		} else {
			b = true
		}
	}
	return b
}

func ConvertUsernameLow(a string) string {
	d := strings.ToLower(a)
	return d
}

func CreateFileSize(mg *mgo.Session) {
	c := mg.DB("file-server").C("sizeLimit")
	mg.Close()

	u := models.SizeLimit{}
	if err := c.Find(bson.M{"_id": ""}).One(&u); err != nil {
		fmt.Println("Error getting SizeLimit", err)
	}

	if u.Limit != 0 {
		MaxSize = u.Limit
	} else {
		b := models.SizeLimit{
			Limit: MaxSize,
		}

		// Write the limitSize to mongo
		if err := c.Insert(b); err != nil {
			fmt.Println("Error: Problem writing Size Limit - ", err)
		}
	}
}
