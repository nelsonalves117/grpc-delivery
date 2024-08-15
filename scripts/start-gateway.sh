#!/bin/sh

curl -d@json/start.json http://localhost:8081/Delivery/Start
curl -d@json/end.json http://localhost:8081/Delivery/End