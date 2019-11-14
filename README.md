# GraphQL Backend

A graphql server written in go for my forum app

## Requirements

- Go
- PostgreSQL

## How to run

```bash
go run *.go
```

## How to build

```bash
go build
```



```sql
CREATE DATABASE forums;
CREATE TABLE threads(
    id serial PRIMARY KEY,
    author varchar(50),
    title varchar(50),
    date_posted TIMESTAMP
);

CREATE TABLE account(
   user_id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (355) UNIQUE NOT NULL,
   created_on TIMESTAMP NOT NULL,
   last_login TIMESTAMP
);

```