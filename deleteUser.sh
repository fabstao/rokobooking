#!/bin/bash

TOKEN="eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiIiwidXNlcm5hbWUiOiJmYWJzIiwicm9sZSI6ImFkbWluIn0sImV4cCI6MTU1MTY1OTg5MSwiaXNzIjoiRmFicyJ9.FpGw5CENntWHGtPVRmKIdBaHWVXt82MFOPMuokjGm41G07OcAWLd_nknzMlPrrKGuMyJXcxlhysjNpyl1wqnRCUuurp4YJSKWLPVf77Cz5BbqhExp3-symS3Nb2rrSlnVeogr5wdqHRK9yNoX7p5G85qEz4xqrmWGBQWn0H_NhIelzRdSGzFtCyGcQjTG-prRJafHK_wceou66261Qw_udbAZhI6GLL7_ajfW6ttMx6ICZqqKcZxcaZ7MY14KFTqYoSUQwVqdrH7vRPRvrDBfZpjmYocn_HUE2pci_tMRWwXuXFPpS1WHqXdeCdaqNUOGEFZSR7xxztvKXzbhTfsdQ"

USER=fabs

curl -XDELETE -H "X-Token: ${TOKEN}" -H "X-Account: ${USER}" http://localhost:8188/user/maintenance7
