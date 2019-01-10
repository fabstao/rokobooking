package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/fabstao/rokobooking/authentication"
	"gitlab.com/fabstao/rokobooking/controllers"
	mgo "gopkg.in/mgo.v2"
)

func getSession(dbhost string, dbname string, dbuser string, dbpasswd string) *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://" + dbuser + ":" + dbpasswd + "@" + dbhost + "/" + dbname)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	// Instantiate a new router
	dbhost := "192.168.0.166"
	dbuser := "roko"
	dbpasswd := "rokoroko"
	dbname := "rokobookdb"
	r := httprouter.New()
	uc := controllers.NewUserController(getSession(dbhost, dbname, dbuser, dbpasswd))
	r.GET("/test", uc.TestAPI)
	r.GET("/user/:username", uc.GetUser)
	r.GET("/users", uc.GetAllUsers)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.GET("/artist/:id", uc.GetArtist)
	r.GET("/artists", uc.GetAllArtists)
	r.POST("/artist", uc.CreateArtist)
	r.POST("/login", authentication.Login)
	r.DELETE("/artist/:id", uc.DeleteArtist)
	http.ListenAndServe("localhost:8189", r)
}
