# Build

GOOS=linux go build
docker build . -t realtime

# Running

Disable ES username and password

$ docker exec -it elasticsearch1 /bin/bash
$ elasticsearch-plugin remove x-pack
$ docker restart elasticsearch1
