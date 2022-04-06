---
exclude: true
---

# Docker Beispiel für hello-world API Service

Beispiel für ein schlankes Docker image mit [Busybox](https://hub.docker.com/_/busybox).

```
$ docker image ls

REPOSITORY             TAG       IMAGE ID       CREATED          SIZE
go-example-container   latest    26b12691b7ae   27 seconds ago   7.72MB
```


## Container Build & Start
```
$ docker build -t go-example-container .
$ docker run -p 8081:8081 go-example-container
```