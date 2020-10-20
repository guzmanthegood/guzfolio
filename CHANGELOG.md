# Changelog

## 0.2.0 (October 20, 2020)

- The following dataloaders have been added and configured
    - UserByID
    - CurrencyByID
    - PortfoliosByUser
- Seeds program has been added, now you can initialize the database with the following command:
    - `go run datastore/seed/seed.go`
- Added references to `README.md`

## 0.1.0 (October 19, 2020)

- First executable version
    - Initial application structure
- `server` package is the main module that runs the GraphQL server
- `datastore` package is the first version of the database connection layer
- Updated `README.md`