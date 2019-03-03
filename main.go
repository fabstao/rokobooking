package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	//dbhost := "192.168.0.166"
	dbhost := "localhost"
	dbuser := "roko"
	dbpasswd := "rokoroko"
	dbname := "rokobookdb"
	r := httprouter.New()
	uc := controllers.NewUserController(getSession(dbhost, dbname, dbuser, dbpasswd))
	r.GET("/test", uc.TestAPI)
	r.GET("/user/:username", uc.GetUser)
	r.GET("/users", uc.GetAllUsers)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:username", uc.DeleteUser)
	r.GET("/artist/:id", uc.GetArtist)
	r.GET("/artists", uc.GetAllArtists)
	r.POST("/artist", uc.CreateArtist)
	r.POST("/login", uc.Login)
	r.DELETE("/artist/:id", uc.DeleteArtist)
	r.POST("/check", uc.CheckT)
	fmt.Println("________________________________________")
	fmt.Println("Listening on 8188")
	fmt.Println("________________________________________")
	http.ListenAndServe("localhost:8188", r)
}
