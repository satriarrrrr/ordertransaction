# Microservice Payment

## Synopsis

This project built for handling order transaction.

## Installation

This project built using Go 1.9.1 as the programming language.


## Initialize Project

### 1. Download Package Dependency

Download all package dependency by execute file dependency.sh:
```
./dependency.sh
```

### 2. Set Configuration File

Duplicate file config.yaml.dist to config.yaml and set each parameter listed:

```
cp config.yaml.dist config.yaml
```

### 3. Build

In folder project, run following command:

```
go build -race
```

There will be binary file named store.

### 4. Database Migration & Seeding

Use file migrate.sql and seed.sql to do database migration and seeding.

### 5. Run

Execute

```
./store -port xxxx -addr "" -config "config.yaml"
```

### 5. System Health Check

After HTTP server is running, you can do several checking to make sure that the service is running well.
> System Health Check (GET)

This endpoint can be used to determine whether service is alive or not.

<http://addr:port/healthcheck>

> Check Connection DB (GET)

This endpoint can be used to determine whether connection to the database is established or not.

<http://addr:port/ping?conn_name=store>


## API Reference

Postman collection: 
Apiary: 

## Tests

> go test -v

## Contributors

Satria Ramadhan <satriaramadhan93@gmail.com>
