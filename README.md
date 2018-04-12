# McFly
Layer to interface between mongo db database and flat files using golang backend.

# Docker
```bash
$ docker network create -d bridge mongo-network
$ docker run --name mongodb -p 27017:27017 -itd --network=mongo-network mongo
$ docker build -t docker/mcfly .
$ docker run --name mcfly -itd --network=mongo-network -p 3000:3000 docker/mcfly
```
