#!/bin/bash
WD=$(dirname $(readlink -f $0))

(
    cd $WD
    echo DOCKER_USER_ID=$(id -u):$(id -g) > .env && echo DOCKER_USER=$USER >> .env
    docker compose up -d
)