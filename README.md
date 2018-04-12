# McFly
Layer to interface between mongo db database and flat files using golang backend.

# Docker
```bash
$ docker run --name mongodb -p 27017:27017 -d mongo
$ docker build -t docker/mcfly .
$ docker run --name mcfly --link mongodb -p 3000:3000 docker/mcfly
```
