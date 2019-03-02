package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/fabstao/rokobooking/authentication"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/fabstao/rokobooking/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var artist = models.Artist{}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// UserController :
type UserController struct {
	session *mgo.Session
}

// NewUserController :
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// TestAPI :
func (uc UserController) TestAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.TestA{
		Name: "Yo",
		Msg:  " Holass",
	}
	uj, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		fmt.Fprintf(w, "{\"ERROR\": \"err\"}")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
	fmt.Println(uj, u)
}

// GetAllUsers :
func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	find := uc.session.DB("rokobookdb").C("users").Find(bson.M{})
	users := find.Iter()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "[")
	size, _ := find.Count()
	i := 0
	for users.Next(&u) {
		i++
		uj, err := json.Marshal(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"ERROR\": \"err\"}")
			return
		}
		if i < size {
			fmt.Fprintf(w, "%s,\n", uj)
		} else {
			fmt.Fprintf(w, "%s", uj)
		}
	}
	fmt.Fprintf(w, "]")
}

// GetUser Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	/*
		fmt.Println(id)
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}
		oid := bson.ObjectIdHex(id)
		fmt.Println(oid)
	*/
	fmt.Println("GET /user")
	fmt.Println(username)
	u := models.User{}

	if err := uc.session.DB("rokobookdb").C("users").Find(bson.M{"username": username}).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Println(uj)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser :
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	uc.session.DB("rokobookdb").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser :
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	fmt.Println("Deleting: ", id)
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB("rokobookdb").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}

// GetAllArtists :
func (uc UserController) GetAllArtists(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Artist{}
	find := uc.session.DB("rokobookdb").C("artists").Find(bson.M{})
	users := find.Iter()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "[")
	size, _ := find.Count()
	i := 0
	for users.Next(&u) {
		i++
		uj, err := json.Marshal(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"ERROR\": \"err\"}")
			return
		}
		if i < size {
			fmt.Fprintf(w, "%s,\n", uj)
		} else {
			fmt.Fprintf(w, "%s", uj)
		}
	}
	fmt.Fprintf(w, "]")
}

// GetArtist Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (uc UserController) GetArtist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	fmt.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)

	fmt.Println("GET /artist")
	u := models.Artist{}

	if err := uc.session.DB("rokobookdb").C("artists").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Println(uj)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateArtist :
func (uc UserController) CreateArtist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Artist{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	uc.session.DB("rokobookdb").C("artists").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteArtist :
func (uc UserController) DeleteArtist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	elh := r.Header.Get("X-Token")
	user := r.Header.Get("X-Account")
	us := models.User{Username: user, Role: "admin"}
	_, err := authentication.ValidateToken(elh, us)
	if err != nil {
		w.WriteHeader(403)
		fmt.Fprintf(w, "{ \"Status\": \"Unauthorized\"  }")
		return
	}
	// Grab id
	id := p.ByName("id")
	fmt.Println("Deleting: ", id)
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB("rokobookdb").C("artists").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}

// CheckT :
func (uc UserController) CheckT(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	elh := r.Header.Get("X-Token")
	//fmt.Println(elh)
	user := r.Header.Get("X-Account")
	us := models.User{Username: user, Role: "admin"}
	fmt.Println(us)
	status, err := authentication.ValidateToken(elh, us)
	fmt.Print("Status: ")
	fmt.Println(status)
	checkError(err)
	salida, err := json.Marshal(status)
	checkError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(salida)
}

// Login :
func (uc UserController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	var userdb models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	checkError(err)

	if err := uc.session.DB("rokobookdb").C("users").Find(bson.M{"username": user.Username}).One(&userdb); err != nil {
		w.WriteHeader(404)
		return
	}
	if user.Username == userdb.Username && user.Passwd == userdb.Passwd {
		user.Passwd = ""
		user.Role = userdb.Role
		fmt.Printf("Successful login: %v  with role: %v \n", user.Username, user.Role)
		token, err := authentication.GenerateJWT(user)
		checkError(err)
		result := models.ResponseToken{Token: token}
		jsonResult, err := json.Marshal(result)
		checkError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		w.Write(jsonResult)
	} else {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Usuario o contraseña inválidos")
		w.Write([]byte("{\"error\":\"Invalid password\"}"))
	}
}
