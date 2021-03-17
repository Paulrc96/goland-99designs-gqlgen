# Go Server

Includes a GraphQL server and a web client.
## Run Server
```sh
go run server.go
```

## Gqlgen
### Generate resolvers and models after a change in `schema.graphqls`
```
go run github.com/99designs/gqlgen generate
```

## Dataloader

### Pre-configuration
Check https://github.com/vektah/dataloaden/issues/35

In the root directory
```
go get github.com/vektah/dataloaden
```

Then
```
mkdir dataloader
cd dataloader
```

Create a file gen.go with the following
```go
package dataloader
```

Then run
```
go run github.com/vektah/dataloaden PostLoader int []*github.com/paul/go-server/graph/model.Post
go run github.com/vektah/dataloaden CommentLoader int []*github.com/paul/go-server/graph/model.Comment
```

## Transactions problem...
https://github.com/lib/pq/issues/81

## URLs
### Graphql Query
POST `http://localhost:8080/query`

Example body:
```graphql
query {
  users (first: 1) {
      id
      name
      last_name
      email
      address
      birthday
      posts {
        title
        description
        comments {
          post_id
          description
       }
     }
  }
}
```

## PostgreSQL
### Increase number of connections
```SQL
alter system set max_connections = 100;
```

## GraphQL Client

github.com/machinebox/graphql