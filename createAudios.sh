#!/bin/bash

curl -XPOST -H "X-Token: $2" -H "X-Account: $1" -H 'Content-Type: application/json' -d '{"name": "Heavy Audio", "rep": "Cisco Alderson", "email": "cisco@haudio.biz", "phone": "3434-555-9812","www": "http://www.haudio.biz", "fb": "facebook.com/HeavyAudio","inst": "@haudio", "twitter": "@haudio", "desc": "Heavy Audio PA, backline","pa": ["Mezcladora 24 canales","Microfonos Shure: 8","Poderes Crown: 2"],"watts": 400.0, "backline":["Ampli guitarra Vox","Bateria Pearl 5 pc platillos Sabian","Amplificador bajo Ampeg"],"luces":["Cerebro DMX","Robot luces LED"]}'  http://localhost:8188/audio

