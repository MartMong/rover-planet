## ROVER PLANET

Requirements
------------

Paddington requires the following to run:

* [Go Programming Language][Go] 1.18+
* [Docker][Docker]

[Go]: https://go.dev
[Docker]: https://www.docker.com/

### Install Dependencies
```sh
go mod download
```

### Run in Local

```sh
go run cmd/rover-planet/main.go  -cli
go run cmd/rover-planet/main.go  -file text.txt
go run cmd/rover-planet/main.go  -rest
```

### Run unit test
```sh
ginkgo -r
```



### example
##### rest
```
curl --location --request POST 'http://localhost:4000/rover' \
--header 'Content-Type: text/plain' \
--data-raw '24
R
F
L
F
L
L
F
R'
```