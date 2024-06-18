#!/bin/sh

/app/migrate -database postgres://$PG_USER:$PG_PASS@$PG_HOST:5432/$PG_DB_NAME?sslmode=disable -path /app/schema/ up
/app/templ generate
/app/library
