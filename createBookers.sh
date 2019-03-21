#!/bin/bash

curl -XPOST -H "X-Token: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiIiwidXNlcm5hbWUiOiJ1c2VyMSIsInJvbGUiOiJ1c2VyIn0sImV4cCI6MTU1MzE2MDU1OCwiaXNzIjoiRmFicyJ9.iixnHxq1voBoMinBUoiSa2fcos4366m68WyINom6rWDmVRspL6wuid4O2LkBZ3_hDcYv7aDT-tLgr6xSyDUZYbyYqxH_gni_QqgxjLPVV6095y1oEvvQ9hlRxcf630zK4YDEhPxTNPlWNf7-2JhpOMcXnVlq7UOUaZPLQIh3wuE-KtK6BkrXopoXhhn5e2tsyLX4vnQbCwtaLmMwwYEtE-UeFYnlvpQTFA5V73367DMDni_oqR2ETpWUEW9PC_cK0mqolSbmQwbQgowcF84V56ql6wLVB5ZimAOJjuBgkpLkHwI-wAx9SbO-rM1zIGw-DPigJ1XKBEZukGqnNVHqkA" \
 -H "X-Account: user1" -H 'Content-Type: application/json' \ 
-d '{"name": "Ramirez Music", "rep": "Brayan Ramirez", "email": "brayan@ramirezmusic.com", "phone": "87277-333-344","www": "http://www.ramirezmusic.com", "fb": "facebook.com/RamirezMusic","inst": "@ramirezmusic", "twitter": "@ramirezmusic", "desc": "Amplia experiencia promoviendo bandas","art":"Music","venues":["Rock Roll Live","Caradura","Vive Latino"],"genres":["Rock", "Salsa","Indie"],"languages":["Español","Inglés"]}' \
 http://localhost:8188/booker
