#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"username": "admin", "passwd": "nosecret", "role": "admin"}' http://localhost:3000/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "fsalaman", "passwd": "nosecret", "role": "superuser"}' http://localhost:3000/user
