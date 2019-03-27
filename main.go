package main

import (
	"fmt"
	"net/http"
	"os"

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

// Crear función que tome variable de entonrno o valor por defecto
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	// Instantiate a new router
	dbhost := getenv("ROKODB_HOST", "localhost")
	dbuser := getenv("ROKODB_USER", "roko")
	dbpasswd := getenv("ROKODB_PASSWD", "rokoroko")
	dbname := getenv("ROKODB_NAME", "rokobookdb")
	PORT := getenv("ROKOPORT", "8188")
	mihost := getenv("ROKOHOST", "0.0.0.0")
	miurl := mihost + ":" + PORT
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
	r.DELETE("/artist/:id", uc.DeleteArtist)
	r.GET("/booker/:id", uc.GetBooker)
	r.GET("/bookers", uc.GetAllBookers)
	r.POST("/booker", uc.CreateBooker)
	r.DELETE("/booker/:id", uc.DeleteBooker)
	r.GET("/venue/:id", uc.GetVenue)
	r.GET("/venues", uc.GetAllVenues)
	r.POST("/venue", uc.CreateVenue)
	r.DELETE("/venue/:id", uc.DeleteVenue)
	r.POST("/login", uc.Login)
	r.POST("/check", uc.CheckT)
	fmt.Println("")
	fmt.Println("________________________________________")
	fmt.Println("")
	fmt.Println(" * ROKOBOOKING * ")
	fmt.Println("")
	fmt.Println("Listening on ", miurl)
	fmt.Println("________________________________________")
	fmt.Println("")
	http.ListenAndServe(miurl, r)
}
