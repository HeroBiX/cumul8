package main

import (
	// Standard library packages
	"html/template"
	"net/http"
	"regexp"

	// Third party packages
	"github.com/HeroBiX/cumul8/server/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

var listLinks string

var templates = template.Must(template.ParseFiles("html/upload.html", "html/main.html"))

type Page struct {
	Status string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func Main(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &Page{Status: controllers.StatusHTML}
	renderTemplate(w, "main", p)
	controllers.StatusHTML = ""
}

func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &Page{Status: controllers.StatusHTML}
	renderTemplate(w, "upload", p)
	controllers.StatusHTML = ""
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

	// Create a new user
	r.POST("/create-user/", uc.CreateUser)

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
