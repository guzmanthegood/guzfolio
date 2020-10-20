# Guzfolio
Guzfolio is a cryptocurrency portfolio tracker with the purpose of using it as a pattern for the implementation of good 
practices in the development of API's in [GraphQL](https://graphql.org) using [Go](https://golang.org). 
As main dependency we will use [gqlgen](https://gqlgen.com) to generate our API, probably the best option.

## Getting Started
1. Download and install [Go 1.13 or greater](https://golang.org/doc/install)
2. Configure Postgres database connection
    - [OPTIONAL] Download and install [Postgres](https://www.postgresqltutorial.com/install-postgresql/)
    - Export (or configure in your favourite IDE) with environment variable `PG_CONNECTION_STRING`  
    `postgres://user:pass@localhost:5432/db_name?sslmode=disable`
3. Initialize database schema and seed with some fake data
    - `go run datastore/seed/seed.go`
4. Start the server running the following command:
    - `go run server/*`

## Authorization
I have chosen a [JSON Web Token (JWT)](https://jwt.io/introduction/) authorization because is a compact and 
self-contained way for securely transmitting information between parties as a JSON object, and they are commonly used. 
In a traditional REST API, when applying the authorization pattern, using a middleware, we can choose which routes 
to secure and which are not, in this way we  can separate the typical register/login calls from the rest of the API 
that we want to secure.

In the case of GraphQL we only have one endpoint, and we cannot use the schema to define the register/login mutations, 
to be honest, in a production environment, the authentication server would be separated in another service, generating 
the tokens that would be used in the service to consume.

I have created two register/login endpoints outside of the GraphQL API context to be able to generate the necessary 
tokens to be able to authenticate. To authenticate with the GraphQL API you have to register as a new user or login 
with an existing user (default pass is _guzfolio1234_) in the next endpoints:

- register new user
    - `/auth/register?email=default@guzfolio.dev&password=guzfolio1234&name=default_name`
- login with existing user
    - `/auth/login?email=user@guzfolio.dev&password=guzfolio1234`

When you obtain your JWT token you can use it in the header of your calls to the GraphQL service with the name 
"Authorization" and value "Bearer xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx". If you are using the Playground you can 
include the following JSON in the HTTP HEADERS section at the bottom.

```json
{
  "Authorization": "Bearer xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

## Dependencies
| **Tech**                                              | **Description**                                                            |
| ----------------------------------------------------  | -------------------------------------------------------------------------- |
| [gqlgen](https://github.com/99designs/gqlgen)         | Go generate based graphql server library                                   |
| [go-chi](https://github.com/go-chi/chi)               | Lightweight, idiomatic and composable router for building Go HTTP services |
| [go-gorm](https://github.com/go-gorm/gorm)            | Fantastic ORM library for Golang, aims to be developer friendly            |
| [dataloaden](https://github.com/vektah/dataloaden)    | Go generate based DataLoader                                               |
| [jwt-go](https://github.com/dgrijalva/jwt-go)         | Golang implementation of JSON Web Tokens (JWT)                             |

## References
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

## License
[MIT](LICENSE)


