# Changelog

## v0.3.0 (October 21, 2020)

- `auth` package has been added
    - login/register routes have been added to register new users and log into the system
    - Authorization middleware has been added to graphQL endpoint
    - Admin functionality added
- GraphQL Schema has been updated
    - profile query has been added (get your own user)
- JWT authorization pattern explained in `README.md`

## v0.2.0 (October 20, 2020)

- The following dataloaders have been added and configured
    - UserByID
    - CurrencyByID
    - PortfoliosByUser
- Seeds program has been added, now you can initialize the database with the following command:
    - `go run datastore/seed/seed.go`
- Added references to `README.md`

## v0.1.0 (October 19, 2020)

- First executable version
    - Initial application structure
- `server` package is the main module that runs the GraphQL server
- `datastore` package is the first version of the database connection layer
- Updated `README.md`