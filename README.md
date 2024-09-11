How to run server locally:
1. Go to project directory: `cd marcopolo`
2. Initialise go module: `go mod init`
3. Tidy up requirements:`go mod tidy`
4. Install the `migrate` CLI tool( Make sure $GOPATH/bin is in your PATH to use the migrate command directly ): `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
5. Run migrations: `migrate -path "./database/migrations" -database "postgres://username:password@127.0.0.1:5432/mydb?sslmode=disable" up`
6. Run app server: `go run cmd/api/main.go`

Go to Postman and simply run the HTTP GET API call:

`http://127.0.0.1:8080/api/users`



A simple Golang webapp built using:
- Gin Web framework
- GORM for DB ORM

Uses:
- Postgres DB

Reference: https://github.com/golang-standards/project-layout?tab=readme-ov-file

```
marcopolo/
├── cmd/
│   ├── api/
│   │   └── main.go           # Entry point for the API server
│   └── cli/
│       └── main.go           # Entry point for the CLI tool
├── internal/
│   ├── handlers/             # Internal HTTP handler functions (organized by domain)
│   │   ├── payment/
│   │   ├── user/
│   │   └── transaction/
│   ├── database/             # Internal database setup and connection logic
│   │   ├── migrations/       # SQL or Go migrations
│   │   ├── seeds/            # Database seed data
│   │   ├── repository/       # GORM-based repository pattern implementations
│   ├── routes/               # Internal route definitions
│   ├── config/               # Internal configuration logic
│   ├── middleware/                # HTTP middleware (e.g., logging, auth)
│   └── services/                  # Internal Business logic services (organized by domain)
│       ├── payment/
│       ├── user/
│       └── transaction/
├── pkg/
│   ├── models/               # Publicly reusable database models
│   └── utils/                # Publicly reusable utility functions
├── test/                          # Centralized testing directory
│   ├── integration/               # Integration tests
│   │   ├── payment_test.go
│   │   ├── user_test.go
│   │   └── transaction_test.go
│   └── e2e/                       # End-to-end tests
│       └── checkout_test.go
└── go.mod                    # Go module file
```