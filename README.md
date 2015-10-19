# Simple Web Server written in Go 
Displays 
 * $WELCOME_MSG
 * Hostname
 * $PID

## To Run:

```
export WELCOME_MSG="Hello World"
go run hello-world.go
```

## To Test:

```
curl http://localhost:80
*** Go - Hello World ! ***
WELCOME_MSG : Hello World
Hostname is : <your-hostname>
Process ID  : 98
```

## Docker

Can be used to create a docker image and run within a container

### To Build
```
docker build -t robertshand/go-hello-world .
```

### To Run
```
docker run -d -p 80:80 -e WELCOME_MSG="Hello World" robertshand/go-hello-world
```

### To Test
```
curl http://$(docker-machine ip default):80
```
Result is
```
*** Go - Hello World ! ***
WELCOME_MSG : Hello World
Hostname is : 9ffe7e421652
Process ID  : 1
```
