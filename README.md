# GOBACK

GOBACK is the backed of Locally written in Golang. 

## PREREQUISITE

1. Install POSTGRESQL, insert dbname, username, password and host in the .env file.(copy .env.example) 
2. Docker

## INSTALLATION

(Without Docker)

1. RUN below command to download necessary packages:

```go
$ go mod download
```

2. RUN below command to run the project:
```go
$ make run
```

(With Docker)

1. Run below command to create and run docker image
```docker
$ docker-compose up
```

2. Check to see if docker is running by listing running docker and if goback container is in the list
```docker
$ docker ps
```

3. To stop docker, RUN
```
docker-compose down
```

4. To enter into docker container, RUN
```docker
docker exec -it goback bash
```