# Food Plan Organizer

The Food Plan Organizer helps people who cook often by creating per-week grocery lists, as well as
creating detailed calory intake overviews.

## Architecture

The application consists of three parts:

1. ETL binary which converts [SR27 nutrition data](http://www.ars.usda.gov/Services/docs.htm?docid=24912) into a [QL](https://github.com/cznic/ql) database.
   The resulting database file is shipped with the application to skip the database creation on the client.
2. A webserver which handles interaction with the SR27 data as well as managing all user input, including recipe management.
3. A single page application written in angularjs, which provides the UI

We're using [atom-shell](https://github.com/atom/atom-shell) to wrap parts 2 & 3 in a executable binary.

## TODOs

- Makefile to generate atom-shell
- non-SR27 data storage
- UI
- Tests

## Commands

Need to be moved into makefile


```
go run cmd/serve/main.go
```