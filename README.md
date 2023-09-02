## 功能


ddd项目模板
```

├─apis
│  ├─graphql
│  ├─grpc
│  └─http
│      ├─app
│      └─h5
├─application
│  ├─events_handler
│  ├─grpc_handler
│  ├─http_handler
│  │  ├─app
│  │  └─h5
│  ├─mq_handler
│  └─timer_handler
├─cmd
│  ├─job
│  │  └─wireinject
│  └─server
│      └─wireinject
├─deploy
│  ├─build
│  └─docker-compose
├─docs
├─domain
│  ├─pay
│  └─user
│      ├─entities
│      ├─events
│      ├─factory
│      ├─interface
│      │  ├─reference
│      │  └─repository
│      ├─service
│      └─value_objects
├─infrastructure
│  ├─berror
│  ├─config
│  ├─db
│  │  ├─connect
│  │  ├─model
│  │  └─repository
│  ├─eventBus
│  ├─http
│  │  ├─middleware
│  │  ├─response
│  │  ├─server
│  │  └─validate
│  ├─lock
│  ├─logx
│  │  ├─formatter
│  │  └─hook
│  ├─redis
│  ├─reference
│  ├─snowflake
│  └─speedLimit
├─mocks
│  ├─repository
│  └─service
├─pkg
│  ├─contextValue
│  ├─contextx
├─scripts
└─test
    └─server
        ├─handler
        ├─repository
        └─service
```
