package models

import "gopkg.in/mgo.v2/bson"

// Profile :  base model for profiles, not DRY
type Profile struct {
	Name    string `json:"name"`
	Rep     string `json:"rep"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Www     string `json:"www"`
	Fb      string `json:"fb"`
	Inst    string `json:"inst"`
	Twitter string `json:"twitter"`
	Desc    string `json:"desc"`
}

// Artist : Model for artists
type Artist struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	*Profile
	Art   string   `json:"art"`
	Genre string   `json:"genre"`
	Rider []string `json:"rider"`
}
