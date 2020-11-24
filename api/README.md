# GraphQL API and a Rest API that works as ETL

## GraphQL

This API is meant to allow mutations and queries to the Dgraph Database. 

It only serves two endpoint localhost:8080/graphql (for queries and mutations) and localhost:8080/admin for configurations 

## How to use 
* In one terminal run:

```
sudo docker run --rm -it -p 8080:8080 -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph dgraph/standalone:v20.03.0
```

* In another terminal and inside /api folder, run: 
```
curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'
```

Now its posible to conect to localhost:8080/graphql, localhost:8080/admin, but also http://localhost:8000/?latest to use Dgraph Ratel,
an interface to easily query and visualize your data.

---

## Rest API

ETL -> Extract, Transform, Load.

This Rest API build in golang and chi extract Customers, Transactions and Products in 3 different formats and load them to the Dgraph database, using the GraphQL API. 


### How to use 

To start the API run from /api:
```
go run main.go
```

Now you can post localhost:3000/v1/load with an optional date to load the data.

If you want to run tests you can use:

```
go test ./...
```
