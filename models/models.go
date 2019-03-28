package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

// Profile :  base model for profiles, not DRY
type Profile struct {
	Uid     bson.ObjectId `json:"uid" bson:"uid"`
	Name    string        `json:"name"`
	Rep     string        `json:"rep"`
	Phone   string        `json:"phone"`
	Email   string        `json:"email"`
	Www     string        `json:"www"`
	Fb      string        `json:"fb"`
	Inst    string        `json:"inst"`
	Twitter string        `json:"twitter"`
	Desc    string        `json:"desc"`
}

// Artist : Model for artists
type Artist struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	*Profile
	Art   string   `json:"art"`
	Genre string   `json:"genre"`
	Rider []string `json:"rider"`
}

// User :
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username"`
	Passwd   string        `json:"passwd,omitempty"`
	Role     string        `json:"role"`
}

// Booker : Model for booking agents
type Booker struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	*Profile
	Venues    []string `json:"venues"`
	Genres    []string `json:"genres"`
	Languages []string `json:"languages"`
}

// Venue : Model for concert venues
type Venue struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	*Profile
	Street   string   `json:"street"`
	Address  string   `json:"address"`
	City     string   `json:"string"`
	ZipCode  string   `json:"zipcode"`
	State    string   `json:"state"`
	Country  string   `json:"country"`
	Location string   `json:"location"`
	Capacity int      `json:"capacity"`
	Audio    []string `json:"audio"`
}

// Audio : Model for audio rental services
type Audio struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	*Profile
	PA       []string `json:"pa"`
	Watts    float32  `json:"watts"` // Watts totales del PA
	Backline []string `json:"backline"`
	Luces    []string `json:"luces"`
}

// TestA :
type TestA struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

// ResponseToken :
type ResponseToken struct {
	Token string `json:"token"`
}

// Claim : From JWT
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
