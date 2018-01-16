## Chat Application

Web application built with Reactjs and Golang using Cassandra as database.

The application has been implemented by following the Udemy course
[Build Realtime Apps](https://www.udemy.com/realtime-apps-with-reactjs-golang-rethinkdb) by @knowthen

Run `go run *.go` for the backend server to start and `webpack-dev-server --port 4001 --hot --inline` for frontend. 
Navigate to `http://localhost:4001`.

It does not offer realtime yet but it would be easy to implement by creating a Cassandra Trigger in Java
and sending the changes to Kafka which then can be consumed in the Golang code.

### Stack

The application uses a number of open source projects to work properly:

- Golang
- Reactjs
- Cassandra
- Kafka : future implementation to offer realtime

### Set Up:

Run `cqlsh -f rtsupport.cql` to create the Keyspace "rtsupport" 
as well as create tables "user", "channel" and "message".