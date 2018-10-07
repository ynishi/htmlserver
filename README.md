# htmlserver [![Build Status](https://travis-ci.org/ynishi/htmlserver.svg?branch=master)](https://travis-ci.org/ynishi/htmlserver)
* simple, plain html server for docker environment
* inspired by json-server https://github.com/typicode/json-server

## Usage
* default port 8080, mount ./html to serve files
```
docker-compose up
# or
docker pull ynishi/htmlserver
docker run --rm -p 8080:8080 -v ${PWD}:/html ynishi/htmlserver
```
* work in container 
```
docker run --rm -it --entrypoint="sh" ynishi/htmlserver 
```
* build and run with raw go bin
```
go build
./htmlserver run
```
