#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"username": "admin", "passwd": "nosecret", "role": "admin"}' http://localhost:3000/user
