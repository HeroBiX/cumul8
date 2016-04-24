package controllers

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// Third Party packages
	"github.com/HeroBiX/cumul8/server/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var StatusHTML string

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func marsjalJson(u models.User) ([]uint8, error) {
	return json.Marshal(u)
}

func WriteContent(w http.ResponseWriter, uj []uint8) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s \n", uj)
}

// Reset Password
func (uc UserController) ResetPwd(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("file-server").C("users")
	session := uc.session.Copy()

	// make sure to close connection to Mongo
	defer session.Close()

	// Grab username
	userUsername := r.FormValue("reset-username")

	// update password
	if err := c.Update(bson.M{"username": userUsername}, bson.M{"$set": bson.M{"password": "password"}}); err != nil {
		w.WriteHeader(404)
		StatusHTML = `Getting User Failed... a hamster died`

		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)
	}
	// new status for login page
	StatusHTML = `Password changed successfully to "password" for: ` + userUsername

	// Redirect back to login
	http.Redirect(w, r, "/", http.StatusFound)

}

// Fetch User Information from db
func GetUser(un string, uc UserController) (models.User, error) {
	c := uc.session.DB("file-server").C("users")
	session := uc.session.Copy()

	// make sure to close connection to Mongo
	defer session.Close()

	// Get users data
	u := models.User{}
	err := c.Find(bson.M{"username": un}).One(&u)

	return u, err
}

// Login
func (uc UserController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab username
	userUsername := r.FormValue("login-username")
	userPassword := r.FormValue("login-password")

	// Get users data
	u, err := GetUser(userUsername, uc)

	// check for error
	if err != nil {
		fmt.Println("error: ", err)
		StatusHTML = `Hamsters Reprot: Wrong Username or Password`
		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// check if username and password is correct
	if userPassword != u.Password {
		StatusHTML = `Hamsters Reprot: Wrong Username or Password`
		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		loggedIn := true
		// uid := u.Id
		err = updateLogin(userUsername, uc, loggedIn)
		if err != nil {
			fmt.Println("error: ", err)
			StatusHTML = `Problem login in, please try again`
			// Redirect back to login
			http.Redirect(w, r, "/", http.StatusFound)
		}
		// Redirect to upload page
		http.Redirect(w, r, "/upload/", http.StatusFound)
	}
}

// CreateUser creates a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()

	// make sure to close connection to Mongo
	defer session.Close()

	fmt.Println("Creating new user")

	// converting form values into variables
	username := r.FormValue("create-username")
	usernamelower := ConvertUsernameLow(username)
	password := r.FormValue("create-password")
	access := r.FormValue("access")

	switch {
	case CheckCreatingUser(username, password, usernamelower, uc) == false:
		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)
	default:
		// Stub an user to be populated from the body
		u := models.User{
			Username:      username,
			UsernameLower: usernamelower,
			Password:      password,
			Access:        access,
		}

		// Populate the user data
		json.NewDecoder(r.Body).Decode(&u)

		// Add an Id
		u.Id = bson.NewObjectId()

		// Write the user to mongo
		uc.session.DB("file-server").C("users").Insert(u)

		// new status for login page
		StatusHTML = `User "` + username + `" was successfully created`

		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)

	}
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()

	// make sure to close connection to Mongo
	defer session.Close()

	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB("file-server").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}

// update login status
func updateLogin(un string, uc UserController, loggedin bool) error {
	c := uc.session.DB("file-server").C("users")
	session := uc.session.Copy()
	var err error

	fmt.Println("Updating the user loggin")

	// make sure to close connection to Mongo
	defer session.Close()

	// check if user is logging in our out
	if loggedin == true {
		// update login status
		err = c.Update(bson.M{"username": un}, bson.M{"$set": bson.M{"Online": "true", "LastLoggedIn": time.Now()}})

	} else {
		// else change status to logged out
		err = c.Update(bson.M{"username": un}, bson.M{"$set": bson.M{"Online": "false"}})
	}

	return err

}
