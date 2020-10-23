# Changelog

## v0.4.0 (October 23, 2020)

### Added
- **Transaction** model has been added
    - A portfolio consists a list of transactions executed at the time of buying or selling cryptocurrencies
    - The objective is to know the difference (profit/lost) of the purchases of these cryptocurrencies
- `seeds.go` now generates transactions with random values for portfolios created by default
    
### Updated
- **Portfolio** model has been updated
    - `totalValue`, `netCost`, `percentChange`, `totalTransactions` aggregated fields has been added
- The concept of fiat currency has been removed from the model. All exchange rates are applied in USD (US Dollar)
- GraphQL queries and mutations has been added to `README.md`

### Fixed
- 1:1 dataloaders has been fixed. The order in which the values were returned was not correct, it's important to 
return the results in the same order that they are requested.
- Removed queries from datastore package that are not used

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