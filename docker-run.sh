#!/bin/sh
docker rm tofso-vv-control
docker run --name tofso-vv-control --env-file .env --network=host --device /dev/gpiochip0 tofso-vv-control
