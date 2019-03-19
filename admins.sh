#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user1", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "user2", "passwd": "nosecret", "role": "user"}' http://localhost:8188/user
