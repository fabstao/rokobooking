#!/bin/bash

curl -XPOST -H "X-Token: $2" -H "X-Account: $1" -H 'Content-Type: application/json' -d '{"name": "Water Blind Live", "rep": "Brayan Perez", "email": "brayan@wblive.biz", "phone": "3434-333-344","www": "http://www.wblive.biz", "fb": "facebook.com/WaterBlindLive","inst": "@wblive", "twitter": "@wblive", "desc": "Rock en vivo, mayores de edad","street":"Gonzalitos 12","address":"Col Garza","city":"Monterrey","zipcode":"23122","state":"Nuevo Leon","country":"Mexico","Location":"Sepa","capacity":80,"audio":["PA","Shure SM57 x 8","Mezcladora Yamaha 24 canales"]}'  http://localhost:8188/venue

