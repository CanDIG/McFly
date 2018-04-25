# McFly
Layer to interface between mongo db database and flat files using golang backend.

# Docker
```bash
$ docker network create -d bridge mongo-network
$ docker run --name mongodb -p 27017:27017 -itd --network=mongo-network mongo
$ docker build -t docker/mcfly .
$ docker run --name mcfly -itd --network=mongo-network -p 3000:3000 docker/mcfly
```

#Usage
```bash
$ make build
$ ./bin/McFly &
```
**Endpoints**
- Go to the endpoint with browser
http://localhost:3000/mcfly/main 
- Upload files
- To to the endpoint with browser for export
http://localhost:3000/mcfly/export 

# Server Depencencies
- Go
- Go Libraries
- MongoDB