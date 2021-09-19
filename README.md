# Customer Project

## Run project 
```bash
go run main.go
```

## Build project
```bash
go build
./customer-project
```

## Run customerimporter tests
```bash
go test ./customerimporter -v
```

## customerimporter package
Package customerimporter reads from the given CSV file and returns a sorted slice of email domains along with the number of customers with e-mail addresses for each domain.

## customeradder package
Package customeradder adds new customers to a given CSV file from fakerapi.
API Docs: https://fakerapi.it/en