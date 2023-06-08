hello:
	echo "hello"

build:
	go build -o .bin/main cmd/main.go

doc:
	swag init -g cmd/main.go --output docs
	#swag init -g cmd/main.go --output docs --parseDependency --parseInternal

run: doc local

t:
	go test -v ./test

local:
	go run cmd/main.go

dev: doc d
qa: doc qa
prod: doc p


d:
	APP_ENV=dev go run cmd/main.go

p:
	APP_ENV=prod go run cmd/main.go

q:
	APP_ENV=qa go run cmd/main.go