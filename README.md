## log-analyzer

This project implements a command-line interface (CLI) program  that analyzes cookie logs in CSV format to identify the most active cookie for a specific date.

**Key Features**

- **Dual Mode:** The program can be run as a CLI tool or as a server, depending on the environment variable set.
- **CSV Parsing:** Handles CSV files containing cookie data and timestamps.
- **Date Filtering:** Filters entries to identify cookies used on a particular date.
- **Most Active Cookie:** Determines the cookie that appears most frequently on the specified date.

**Project Structure**

The project directory is organized as follows:
```
.
├── README.md
├── cmd
│   ├── cli.go
│   ├── main.go
│   └── server.go
├── cookie_log.csv
├── go.mod
├── go.sum
└── internal
    ├── config
    ├── constants
    │         └── time.go
    ├── controllers
    ├── handlers
    │         ├── log_extractor.go
    │         └── log_extractor_test.go
    ├── logger
    │         └── logger.go
    ├── models
    ├── repositories
    └── services
```

**Running the Project**

**CLI Mode:**

1. Set the `ENV` environment variable to `cli`.
2. Run `go run cmd/main.go cmd/cli.go cmd/server.go -f <csv_file_path> -d <date>`, replacing `<csv_file_path>` with the path to your CSV file and `<date>` with the date in YYYY-MM-DD format.

**Server Mode:** [TODO]

1. Set the `ENV` environment variable to `server`.
2. Run `go run cmd/main.go cmd/cli.go cmd/server.go`.
3. Send a GET request to `http://localhost:8080/most-active-cookie?date=<date>`, replacing `<date>` with the date in YYYY-MM-DD format. The response will contain the most active cookie for the specified date.

**Dependencies**

This project uses the Go programming language and requires the following dependencies:

* You can install them using the `go mod download` command.

**Further Development**

This project serves as a foundation for building a more comprehensive log analysis system. Potential areas for future development include:

- [ ] Support for different log file formats
- [ ] Advanced filtering capabilities
- [ ] Run it as a server

