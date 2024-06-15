#!/usr/bin/make

IN_APP=docker-compose exec library

run: generate run_app

generate:
	$(IN_APP) /app/templ generate

run_app:
	$(IN_APP) /app/library

up:
	docker-compose up

up_build:
	docker-compose up --build
