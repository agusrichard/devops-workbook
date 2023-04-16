#!/bin/sh

docker stop jenkins-blueocean
docker rm jenkins-blueocean

docker stop alpine-socat
docker rm alpine-socat