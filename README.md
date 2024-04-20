# GO API SERVER
Template Golang for Backend using Clean Architecture

# Project Tree
```bash
├── Dockerfile
├── README.md
├── adapters
│   ├── mappers
│   │   └── userMaps.go
│   ├── orm
│   │   └── usersOrm.go
│   └── repository
│       └── usersRepository.go
├── api
│   ├── dto
│   │   └── userDto.go
│   ├── handlers
│   │   └── v1
│   │       └── userHandler.go
│   ├── presenter
│   │   └── general.go
│   └── routers
│       └── userRouter.go
├── cmd
│   ├── webserver
│   │   └── main.go
│   └── worker
│       └── main.go
├── config
│   ├── database.go
│   └── environtment.go
├── domain
│   ├── entities
│   │   └── userModel.go
│   └── services
│       └── userService.go
├── go.mod
├── go.sum
└── pkg
    ├── logger
    │   └── logger.go
    ├── middleware
    │   ├── cors.go
    │   ├── error_handler.go
    │   └── request_logger.go
    └── util
        └── encrypt.go

```

### Library Using
- **Fiber v2**: Framework.
- **Viper**: Management configuration file.
- **Logrus**: Logger.
- **GORM**: ORM library for Go.