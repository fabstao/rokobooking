#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance1", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance2", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance3", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance4", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance5", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance6", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance7", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "maintenance8", "passwd": "nosecret", "role": "maintenance"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "admin", "passwd": "nosecret", "role": "admin"}' http://localhost:8188/user
curl -XPOST -H 'Content-Type: application/json' -d '{"username": "fabs", "passwd": "yomeroles", "role": "admin"}' http://localhost:8188/user
