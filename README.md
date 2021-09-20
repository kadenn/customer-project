# Customer Project

## Features
- CSV reading and email extraction(import customer)
- can import 1m+ customers in 1 second
- used only Go Standard Library
- adds new customers to dataset from [fakerapi](https://fakerapi.it/en) concurrently using Goroutines

## Run project 
```bash
go run main.go
```

## Build project
```bash
go build
./customer-project
```

## customerimporter package
Package customerimporter reads from the given CSV file and returns a sorted slice of email domains along with the number of customers with e-mail addresses for each domain.

## Test customerimporter
```bash
go test ./customerimporter -v
```

## customeradder package
Package customeradder adds new customers to a given CSV file from fakerapi.
API Docs: https://fakerapi.it/en
