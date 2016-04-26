package controllers

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
func GetUser(un string, mg *mgo.Session) (models.User, error) {
	c := mg.DB("file-server").C("users")
	session := mg.Copy()
	defer session.Close()

	// Get users data
	u := models.User{}
	err := c.Find(bson.M{"username": un}).One(&u)

	return u, err
}

var CurrentUser string

func PleaseLogin(w http.ResponseWriter, r *http.Request) {
	StatusHTML = `Please Login`
	// Redirect back to login
	http.Redirect(w, r, "/", http.StatusFound)
}

// Login
func (uc UserController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()
	defer session.Close()

	// Grab username
	userUsername := r.FormValue("login-username")
	userPassword := r.FormValue("login-password")

	mg := uc.session
	// Get users data
	u, err := GetUser(userUsername, mg)

	// check for error
	if err != nil {
		fmt.Println("error: ", err)
		StatusHTML = `Hamster Reports: Wrong Username or Password`
		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// check if username and password is correct
	if userPassword != u.Password {
		StatusHTML = `Hamster Reports: Wrong Username or Password`
		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		CurrentUser = userUsername
		// Redirect to upload page
		http.Redirect(w, r, "/upload/", http.StatusFound)
	}
}

var MaxSize int64 = 5

// Upload file
func (uc UserController) Upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()
	defer session.Close()

	if r.ContentLength/1000000 > MaxSize { // making sure it counts in MB
		http.Error(w, "File is to big... please limit yourself", http.StatusExpectationFailed)
		return

	} else {

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)

		// save file to server
		f, err := os.OpenFile("./data/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)

		// add filename to DB
		mgo := uc.session
		err = AddFileName(mgo, handler.Filename)
		if err != nil {
			fmt.Println("Add filename error: ", err)
			StatusHTML = `Problem uploading filename to DB`
			http.Redirect(w, r, "/upload/", http.StatusFound)
			return

		} else {
			StatusHTML = `File uploaded successfully`
			http.Redirect(w, r, "/upload/", http.StatusFound)
		}
	}
}

// Download File
func (uc UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()
	defer session.Close()

	url := "./data/"
	fileName := r.FormValue("bobsFile")

	f, err := ioutil.ReadFile(url + fileName)
	if err != nil {
		fmt.Println("read file", err)
		return
	}
	// defer f.Close()

	err = ioutil.WriteFile("./download/"+fileName, f, 0666)
	if err != nil {
		fmt.Println("write file", err)
		return
	}

	fmt.Println("DOWNLOADING FILE!")
	StatusHTML = `File downloaded successfully`
	// Redirect to upload page
	http.Redirect(w, r, "/upload/", http.StatusFound)
}

// change the file size limit
func (uc UserController) Limitsize(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("file-server").C("users")
	session := uc.session.Copy()
	defer session.Close()

	u := models.User{}
	if err := c.Find(bson.M{"username": CurrentUser}).One(&u); err != nil {
		StatusHTML = `The unicorns was able to log you in but not to list your files`
		fmt.Println("Error listing users files", err)
	}

	if u.Access != "admin" {
		StatusHTML = `You are do not have admin rights...`
		// Redirect to upload page
		http.Redirect(w, r, "/upload/", http.StatusFound)

	} else {

		var newSize int64
		a := r.FormValue("size-limit")
		newSize, err := strconv.ParseInt(a, 16, 32)
		if err != nil {
			StatusHTML = `Please insert a number`
			// Redirect to upload page
			http.Redirect(w, r, "/upload/", http.StatusFound)

		} else {

			mg := uc.session
			err := ChangeSizeRestriction(newSize, mg)
			if err != nil {
				fmt.Println("Error: Problem updating MaxSize - ", err)

			} else {
				MaxSize = newSize
				StatusHTML = `file size limit has been changed`
				// Redirect to upload page
				http.Redirect(w, r, "/upload/", http.StatusFound)
			}
		}
	}
}

func ChangeSizeRestriction(s int64, mg *mgo.Session) error {

	// Update size limit to the DB
	err := mg.DB("file-server").C("limit").Update(bson.M{"limit": ""}, bson.M{"$set": bson.M{"limit": s}})

	return err
}

// Get list of all users files
func ListFiles(mg *mgo.Session) string {
	c := mg.DB("file-server").C("users")
	session := mg.Copy()
	defer session.Close()

	// Get users data
	u := models.User{}
	if err := c.Find(bson.M{"username": CurrentUser}).One(&u); err != nil {
		StatusHTML = `The unicorns was able to log you in but not to list your files`
		fmt.Println("Error listing users files", err)
	}

	dbList := u.Filename
	var list string

	for _, i := range dbList {
		list += creatingHTMLcode(i)
	}
	return list
}

func creatingHTMLcode(bob string) string {
	tempString := `<p><form action="/get/" method="GET">` + bob + ` <input type="hidden" name="bobsFile"value="` + bob + `" /> <input type="submit" value="Download"></form></p>`

	return tempString
}

// adding filenames to DB
func AddFileName(mg *mgo.Session, fn string) error {
	c := mg.DB("file-server").C("users")
	session := mg.Copy()
	defer session.Close()

	u := models.User{}
	if err := c.Find(bson.M{"username": CurrentUser}).One(&u); err != nil {
		StatusHTML = `The unicorns was able to log you in but not to list your files`
		fmt.Println("Error listing users files", err)
	}

	// Grab username
	userUsername := CurrentUser
	u.Filename = append(u.Filename, fn)

	// add filename to DB
	err := c.Update(bson.M{"username": userUsername}, bson.M{"$set": bson.M{"filename": u.Filename}})

	return err
}

// CreateUser creates a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()
	defer session.Close()

	// converting form values into variables
	username := r.FormValue("create-username")
	usernamelower := ConvertUsernameLow(username)
	password := r.FormValue("create-password")
	access := r.FormValue("access")

	if CheckCreatingUser(username, password, usernamelower, uc) == false {
		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// Stub an user to be populated from the body
		u := models.User{
			Username:      username,
			UsernameLower: usernamelower,
			Password:      password,
			Access:        access,
		}

		// Populate the user data
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			fmt.Println("Error: Problem populating the new user data - ", err)
		}

		// Add an Id
		u.Id = bson.NewObjectId()

		// Write the user to mongo
		if err := uc.session.DB("file-server").C("users").Insert(u); err != nil {
			fmt.Println("Error: Problem populating the new user data - ", err)
		}

		// new status for login page
		StatusHTML = `User "` + username + `" was successfully created`

		// Redirect back to login
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func FileSizeDB(mg *mgo.Session) {
	c := mg.DB("file-server").C("sizeLimit")
	session := mg.Copy()
	defer session.Close()

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
