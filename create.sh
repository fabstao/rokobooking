#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"username": "operator", "passwd": "nosecret", "role": "operator"}' http://localhost:3000/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "boss", "passwd": "nosecret", "role": "admin"}' http://localhost:3000/user
