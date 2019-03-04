#!/bin/bash

curl -XPOST -d "{\"username\": \"$1\", \"passwd\": \"$2\"}" http://localhost:8188/login