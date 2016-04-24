package controllers

import (
	// Standard Library packages
	"regexp"
	"strings"

	// Thirt Party packages
	"github.com/HeroBiX/cumul8/server/models"
	"gopkg.in/mgo.v2/bson"
)

var b bool

func CheckCreatingUser(un, pass, lun string, uc UserController) bool {

	// See if username looks good
	switch {
	case EnoughCharacters(un, 2) == false:
		StatusHTML = "I do think you need more characters in your username... just a feeling"
		b = false

	case NoFunkyCharacters(un) == false:
		StatusHTML = "The unicorns sense funky characters in your UserName"
		b = false

	case ExistingUserName(un, lun, uc) == false:
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
func ExistingUserName(un, lun string, uc UserController) bool {
	u := models.User{}

	// find if the username exist
	// Figure out a good way to handle errors here
	uc.session.DB("file-server").C("users").Find(bson.M{"username": un}).One(&u)

	if un == u.Username {
		b = false
	} else {

		// find is lower case username exist
		uc.session.DB("file-server").C("users").Find(bson.M{"usernamelower": lun}).One(&u)

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
