# Guzfolio

Guzfolio is a cryptocurrency portfolio tracker with the purpose of using it as a pattern for the implementation of good 
practices in the development of API's in [GraphQL](https://graphql.org) using [Go](https://golang.org). 
As main dependency we will use [gqlgen](https://gqlgen.com) to generate our API, probably the best option.

We will initially follow the example [gqlgen-pg-todo-example](https://github.com/oshalygin/gqlgen-pg-todo-example) 
developed by [oshalygin](https://github.com/oshalygin) (thanks for your work!) to later add new more complex 
functionality like as authorization, nested elements in the data model, unit test, etc...  

# Getting Started
1. Download and install [Go 1.13 or greater](https://golang.org/doc/install)
2. Configure Postgres database connection
    - [OPTIONAL] Download and install [Postgres](https://www.postgresqltutorial.com/install-postgresql/)
    - Export (or configure in your favourite IDE) the environment variable `PG_CONNECTION_STRING`  
    `postgres://user:pass@localhost:5432/db_name?sslmode=disable`
3. Start the server running the following command:
    - `go run server/*`


# Dependencies

| **Tech**                                      | **Description**                                                            |
| --------------------------------------------- | -------------------------------------------------------------------------- |
| [gqlgen](https://github.com/99designs/gqlgen) | go generate based graphql server library                                   |
| [go-chi](https://github.com/go-chi/chi)       | lightweight, idiomatic and composable router for building Go HTTP services |
| [go-gorm](https://github.com/go-gorm/gorm)    | fantastic ORM library for Golang, aims to be developer friendly            |

# License

[MIT](LICENSE)


