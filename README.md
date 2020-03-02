# D2H CLI Application

To run the application:

1. Create two MySQL DBs name with name `d2h` & `d2h-test`
2. Import `d2h.sql` file to both the DBs.
3. Once above steps are done, navigate to the project directory.
4. Run the code using `go run main.go`

### Tree
```shell
.
├── d2h.sql
├── db
│   └── db.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
├── services
    ├── cmd
    │   └── service
    │       └── service.go
    └── pkg
        ├── handlers
        │   └── handlers.go
        ├── menu
        │   └── menu.go
        ├── models
        │   ├── channels.go
        │   ├── info.go
        │   ├── packs.go
        │   ├── services.go
        │   ├── subscriptions.go
        │   └── users.go
        ├── repositories
        │   ├── repositories.go
        │   └── repositories_test.go
        └── service
            └── service.go

```