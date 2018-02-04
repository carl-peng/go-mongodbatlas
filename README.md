# go-mongodbatlas [![Build Status](https://travis-ci.org/akshaykarle/go-mongodbatlas.svg?branch=master)](https://travis-ci.org/akshaykarle/go-mongodbatlas) [![GoDoc](https://godoc.org/github.com/akshaykarle/go-mongodbatlas/mongodb?status.png)](https://godoc.org/github.com/akshaykarle/go-mongodbatlas/mongodb)
A Go client library for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

## Getting started
MongoDB Atlas uses the [Digest Access Authentication](https://tools.ietf.org/html/rfc2069) and doesn't support Basic Auth. Follow the example in [example.go](example.go). We create a httpClient using [go-http-digest-auth-client](https://github.com/xinsnake/go-http-digest-auth-client) and pass it over to mongodb. To run the example, just run:
```
go run examples/clusters.go <username> <mongodb-atlas-api-key> <group-id>
```

## TODO
* Add support for Monitoring & Logging
