# http-boilerplate

An example of a web server in Go.

## Docker

Build image

```
docker build -t go-docker .
```

Image list

```
docker image ls
```

Remove image

```
docker image rm <image id>
```

Run docker image

```
docker run -d -p 8080:8080 http-boilerpalate
```

Container list

```
docker container ls
```

Stop container

```
docker container stop <container id>
```
