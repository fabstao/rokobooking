package main

import (
	"fmt"
	"net/http"
	"os"

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

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
	// Instantiate a new router
	dbhost := getenv("ROKODB_HOST","localhost")
	dbuser := getenv("ROKODB_USER","roko")
	dbpasswd := getenv("ROKODB_PASSWD","rokoroko")
	dbname := getenv("ROKODB_NAME","rokobookdb")
	PORT := getenv("ROKOPORT","8188")
	mihost := getenv("ROKOHOST","0.0.0.0")
	miurl := mihost+":"+PORT
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
	r.POST("/check", uc.CheckT)
	fmt.Println("________________________________________")
	fmt.Println("Listening on ",miurl)
	fmt.Println("________________________________________")
	http.ListenAndServe(miurl, r)
}
