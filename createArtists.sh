#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"name": "VALTIS", "rep": "Fabs Saldom", "email": "fabs@valtis.mx", "phone": "555-53426","www": "http://www.valtis.mx", "fb": "facebook.com/Valtis.Oficial","inst": "@valtis_oficial", "twitter": "@valtis_oficial", "desc": "VALTIS es una banda de rock mexicana que busca su propio sonido dentro del rock analogo y melodico","art":"Music","genre":"Rock","rider":["backline: Vox, Marshall, MOTUS", "drumkit: Tama 6pc, Zildjian A: Ride, 2xCrash, Splash, China, Hihats","mics: Shure SM58"]}' http://localhost:3000/artist
