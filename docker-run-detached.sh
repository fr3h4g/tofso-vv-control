#!/bin/sh
docker rm tofso-vv-control
docker run --name tofso-vv-control -d --network=host --env-file .env --device /dev/gpiochip0 tofso-vv-control
