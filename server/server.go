package main

import (
	// Standard library packages
	"html/template"
	"net/http"

	// Third party packages
	"github.com/HeroBiX/cumul8/server/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

var listLinks string

var templates = template.Must(template.ParseFiles("html/upload.html", "html/main.html"))

type Page struct {
	Status    string
	User      string
	FileSize  int64
	ListFiles template.HTML
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Main(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &Page{Status: controllers.StatusHTML}
	renderTemplate(w, "main", p)
	controllers.StatusHTML = ""
	controllers.CurrentUser = ""
}

func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if controllers.CurrentUser == "" {
		controllers.PleaseLogin(w, r)
	} else {
		uc := controllers.NewUserController(getSession())
		y := template.HTML(controllers.ListFiles(uc))
		p := &Page{Status: controllers.StatusHTML, User: controllers.CurrentUser, ListFiles: y, FileSize: controllers.MaxSize}
		renderTemplate(w, "upload", p)
		controllers.StatusHTML = ""
	}
}

func main() {

	// Instantiate a new router
	r := httprouter.New()

	r.GET("/", Main)
	r.GET("/main/", Main)
	r.GET("/upload/", Upload)

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Reset Password
	r.POST("/resetpwd/", uc.ResetPwd)

	// Login
	r.POST("/login/", uc.Login)

	// Upload file
	r.POST("/uploadfiles/", uc.Upload)

	// Create a new user
	r.POST("/create-user/", uc.CreateUser)

	// Download file
	r.GET("/get/", uc.Get)

	// Change file size
	r.POST("/limitsize/", uc.Limitsize)

	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)

	http.ListenAndServe(":8080", r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
