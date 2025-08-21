# Simple Rest Api with GO

This project is a simple REST API built with **Go** that demonstrates:

### Prerequisites
- Go 1.23.11
- [Driver Mysql](https://github.com/go-sql-driver/mysql) for connecting database
- [Validator](https://github.com/go-playground/validator) for validate request
- [Http Router](https://github.com/julienschmidt/httprouter) for manage router
- [Google Wire](https://github.com/google/wire) for generate injector
- [Testify](https://github.com/stretchr/testify) for help me write unit test
- [Logrus](https://github.com/sirupsen/logrus) for manage log
- [Lumberjack](https://github.com/natefinch/lumberjack) for manage rolling file log
- [Google UUID](https://github.com/google/uuid) for generate uuid

### Install dependencies
```bash
go mod tidy
```

### Generate Injector
make sure injector in 'injector.go'
```bash
wire
```

### Run application
```bash
go run main.go
```

### Run test
```bash
go test -v ./...
```



