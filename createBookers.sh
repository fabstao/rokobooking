#!/bin/bash

curl -XPOST -H "X-Token: $2" -H "X-Account: $1" -H 'Content-Type: application/json' -d '{"name": "Ramirez Music", "rep": "Brayan Ramirez", "email": "brayan@ramirezmusic.com", "phone": "87277-333-344","www": "http://www.ramirezmusic.com", "fb": "facebook.com/RamirezMusic","inst": "@ramirezmusic", "twitter": "@ramirezmusic", "desc": "Amplia experiencia promoviendo bandas","art":"Music","venues":["Rock Roll Live","Caradura","Vive Latino"],"genres":["Rock", "Salsa","Indie"],"languages":["Español","Inglés"]}' http://localhost:8188/booker 
 
