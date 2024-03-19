## log-analyzer

This project implements a command-line interface (CLI) program  that analyzes cookie logs in CSV format to identify the most active cookie for a specific date.

**Project Structure**

The project directory is organized as follows:
```
├── Dockerfile 
├── Makefile                    
├── README.md
├── cmd/
│   ├── cli.go
│   ├── main.go
│   └── server.go
├── cookie_log.csv
├── cookie_log1.csv
├── directory_tree.md
├── go.mod
├── go.sum
├── internal/
│   ├── config/
│   ├── constants/
│   │   ├── server.go               (Opt: Server related Constants)
│   │   ├── time.go                 (Time related Constants)
│   │   └── utils.go                (Other Utils)
│   ├── controllers/
│   │   ├── log_extractor.go        (Opt: Controller Layer for Server Impl)
│   │   └── log_extractor_test.go
│   ├── factories/
│   │   └── log_extractor.go        (Opt:Factory for creating log controller and services for server Impl)
│   ├── handlers/
│   │   ├── log_extractor.go        (Business logic to extact most active cookie)
│   │   └── log_extractor_test.go
│   ├── logger/
│   │   └── logger.go               (Project log config)
│   ├── models/
│   │   └── log.go                  (Opt: Model for log server impl)
│   ├── repositories/
│   ├── router/
│   │   └── router.go               (Opt: router for server Impl)  
│   └── services/
│       ├── log_extractor.go        (Opt: Service layer for server Impl) 
│       └── log_extractor_test.go
├── sample.env
```

**Running the Project**

**Using Makefile**
1. Use `make install` to install the dependencies
2. Use `make build` to build the `most_active_cookie` binary 
3. Run `./most_active_cookie cookie_log.csv -d 2018-12-09` to directly run
3. Opt: Use `make run_cli ARGS="<csv_file_path> -d <date>`, replacing `<csv_file_path>` with the path to your CSV file and `<date>` with the date in YYYY-MM-DD format.

**Developing locally**
1. Use `make install` to install the dependencies
2. Execute `go run cmd/main.go cmd/cli.go cookie_log.csv -d 2018-12-09` to run the project without building

**Run test**
1. Use `make test` to run the testcases
2. Use `make coverage` to check the coverage for the files
