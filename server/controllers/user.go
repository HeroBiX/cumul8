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
		err = updateLogin(userUsername, uc, loggedIn)
		if err != nil {
			fmt.Println("error: ", err)
			StatusHTML = `Problem login in, please try again`
			http.Redirect(w, r, "/", http.StatusFound)
		}
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

	fmt.Println("size ", r.ContentLength)

	if r.ContentLength/1000000 > MaxSize { // making sure it counts in MB
		http.Error(w, "File is to big... please limit yourself", http.StatusExpectationFailed)
		return
		// StatusHTML = `File is to big... please limit yourself to: ` + strconv.FormatInt(MaxSize, 10) + `MB`
		// http.Redirect(w, r, "/upload/", http.StatusFound)
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
		err = addFileName(uc, handler.Filename)
		if err != nil {
			fmt.Println(err)
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
			MaxSize = newSize

			StatusHTML = `file size limit has been changed`
			// Redirect to upload page
			http.Redirect(w, r, "/upload/", http.StatusFound)
		}
	}
}

// Get list of all users files
func ListFiles(uc *UserController) string {
	c := uc.session.DB("file-server").C("users")
	session := uc.session.Copy()
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
func addFileName(uc UserController, fn string) error {
	c := uc.session.DB("file-server").C("users")
	session := uc.session.Copy()
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
	defer session.Close()

	var err error

	fmt.Println("Updating the user loggin")

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
