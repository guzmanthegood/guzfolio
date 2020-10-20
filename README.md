# Guzfolio

Guzfolio is a cryptocurrency portfolio tracker with the purpose of using it as a pattern for the implementation of good 
practices in the development of API's in [GraphQL](https://graphql.org) using [Go](https://golang.org). 
As main dependency we will use [gqlgen](https://gqlgen.com) to generate our API, probably the best option.

# Getting Started
1. Download and install [Go 1.13 or greater](https://golang.org/doc/install)
2. Configure Postgres database connection
    - [OPTIONAL] Download and install [Postgres](https://www.postgresqltutorial.com/install-postgresql/)
    - Export (or configure in your favourite IDE) with environment variable `PG_CONNECTION_STRING`  
    `postgres://user:pass@localhost:5432/db_name?sslmode=disable`
3. Initialize database schema and seed with some fake data
    - `go run datastore/seed/seed.go`
4. Start the server running the following command:
    - `go run server/*`


# Dependencies

| **Tech**                                              | **Description**                                                            |
| ----------------------------------------------------  | -------------------------------------------------------------------------- |
| [gqlgen](https://github.com/99designs/gqlgen)         | go generate based graphql server library                                   |
| [go-chi](https://github.com/go-chi/chi)               | lightweight, idiomatic and composable router for building Go HTTP services |
| [go-gorm](https://github.com/go-gorm/gorm)            | fantastic ORM library for Golang, aims to be developer friendly            |
| [dataloaden](https://github.com/vektah/dataloaden)    | go generate based DataLoader                                               |

# References

- A simple, no fuss, example thats updated regularly to stay current with the API landscape
    - [github.com/oshalygin/gqlgen-pg-todo-example](https://github.com/oshalygin/gqlgen-pg-todo-example)
- Youtube Golang Grapqhql + gqlgen tutorial
    - [Golang GraphQL Tutorial Youtube](https://youtu.be/A6lDNao00WQ)
    - [github.com/EQuimper/youtube-golang-graphql-tutorial](https://github.com/EQuimper/youtube-golang-graphql-tutorial)
- Other resources
    - [Dive into GraphQL](https://medium.com/@ivan.corrales.solera/dive-into-graphql-9bfedf22e1a)
    - [Create Your First Simple GraphQL Golang Application with GQLGen](https://medium.com/@ktrilaksono/create-your-first-simple-graphql-golang-application-with-go-gqlgen-793e11dc5fc4)
    - [How To Create GraphQL Server With Golang](https://dev.to/glyphack/introduction-to-graphql-server-with-golang-1npk)
    - [Introducing gqlgen: a GraphQL Server Generator for Go](https://99designs.com.au/blog/engineering/gqlgen-a-graphql-server-generator-for-go/)

# License

[MIT](LICENSE)


