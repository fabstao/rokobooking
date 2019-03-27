#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user1", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user2", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user3", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user4", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user5", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user6", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user7", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user8", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
