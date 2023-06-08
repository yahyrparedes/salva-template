# Salva Template (Go Architecture)

Salva template easily starter for development apis in go. 

# Dependency

1. [x] Viper (https://github.com/spf13/viper)
2. [x] Gorm (https://github.com/go-gorm/gorm)
3. [x] Fiver (https://github.com/gofiber/fiber)
4. [x] Yaml (https://github.com/go-yaml/yaml/tree/v2.4.0)
5. [x] Salva (https://github.com/yahyrparedes/salva)


## How to run local:

```bash
make run
o
swag init -g cmd/main.go --output docs
go run cmd/main.go
```

## How to run prod:

```bash
make prod
o
swag init -g cmd/main.go --output docs
APP_ENV=prod go run cmd/main.go
```



## Generate docs 
https://github.com/swaggo/swag

```bash
swag init -g cmd/main.go --output docs
```


## Salva CLI

- Add new endpoint for app 
```bash
salva magic 
```


> Example endpoint User
```bash
salva magic  user
```

