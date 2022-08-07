# Transaction App
Transaction app is a transaction generation and display application. It provides a RestApi and web application user interface for displaying
and interacting with transactions. Implemented with Golang and MongoDB due to their ability to easily scale in a distributed system.

## Installing and Running

### Using docker
After installing [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/)
```
$ cd transaction
$ docker-compose build
$ docker-compose up
```

#### Run
Open url in browser `http://localhost:3000/`

#### Testing
```
$ docker exec -it webapi bash
$ go test
```


#### Author
Konrad Djimeli


