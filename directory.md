# ディレクトリとファイル構造

```
cmd
   main.go
pkg
  adapter
   database
      config.go
      database.go
   repository
      model.go
      repository.go
   web
      server.go
     controller
        controller.go
     routing
        common.go
        routing_develop.go
        routing_production.go
     template
        template.go
  app
   config
       config.go
   dependencies
       dependencies.go
  common
   constant
      constant.go
   helper
     repository
         repository.go
  entity
     entity.go
  infrastructure
   dependencies
      init_develop.go
      init_production.go
   repository
     database.go
     repository.go
    model
        model.go
   web
       web.go
  logic
     dependencies.go
     logics.go
     types.go
    dependencies
       dependencies.go
    sales
       logic.go
    users
        logic.go
```

```
.
├── cmd/
│   └── main.go
└── pkg/
    ├── adapter/
    │   ├── database/
    │   │   ├── config.go
    │   │   └── database.go
    │   ├── repository/
    │   │   ├── model.go
    │   │   └── repository.go
    │   └── web/
    │       ├── server.go
    │       ├── controller/
    │       │   └── controller.go
    │       ├── routing/
    │       │   ├── common.go
    │       │   ├── routing_develop.go
    │       │   └── routing_production.go
    │       └── template/
    │           └── template.go
    ├── app/
    │   ├── config/
    │   │   └── config.go
    │   └── dependencies/
    │       └── dependencies.go
    ├── common/
    │   ├── constant/
    │   │   └── constant.go
    │   └── helper/
    │       └── repository/
    │           └── repository.go
    ├── entity/
    │   └── entity.go
    ├── infrastructure/
    │   ├── dependencies/
    │   │   ├── init_develop.go
    │   │   └── init_production.go
    │   ├── repository/
    │   │   ├── database.go
    │   │   ├── repository.go
    │   │   └── model/
    │   │       └── model.go
    │   └── web/
    │       └── web.go
    └── logic/
        ├── dependencies.go
        ├── logics.go
        ├── types.go
        ├── dependencies/
        │   └── dependencies.go
        ├── sales/
        │   └── logic.go
        └── users/
            └── logic.go
```
