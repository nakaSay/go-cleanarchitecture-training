# go-cleanarchitecture-training

Example of a Go application using a clean architecture.

## Install

Preparing Sources
```
cd $WORKDIR
git clone https://github.com/nakaSay/go-cleanarchitecture-training.git
```

## Launching Applications
Running with Docker
```
cd $WORKDIR/go-cleanarchitecture-training/docker
docker-compose up
```

Start API Server
```
docker container exec -it docker_server_1  bash
make run
```

Confirmation of datastore using gcloud-gui
```
http://localhost:8082
```

## Application Features
```
// Save post
curl -XPOST localhost:8080/createPost -H "Content-Type: application/x-www-form-urlencoded" -d "text=hoge"

// Get all posts
curl -X GET localhost:8080/posts

// Update post
curl -XPOST localhost:8080/updatePost -H "Content-Type: application/x-www-form-urlencoded" -d "id=○○&text=~~~~~~"

//Delete post
curl -X GET "localhost:8080/removePost?id=○○"

```
