Reference: https://github.com/golang-standards/project-layout?tab=readme-ov-file

myapp/
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