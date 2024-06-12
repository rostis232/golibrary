run:
	/home/rpylypiv/go/bin/templ generate
	go run cmd/web/main.go

generate:
	templ generate

up:
	docker-compose up

up_build:
	docker-compose up --build
