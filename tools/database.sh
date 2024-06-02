#!/bin/bash

source .env
schema=$(echo $DATABASE_SCHEMA | sed "s/\r//")
user=$(echo $DATABASE_USER | sed "s/\r//")
password=$(echo $DATABASE_PASSWORD | sed "s/\r//")
host=$(echo $DATABASE_HOST | sed "s/\r//")
port=$(echo $DATABASE_PORT | sed "s/\r//")
name=$(echo $DATABASE_NAME | sed "s/\r//")
ssl_mode=$(echo $DATABASE_SSL_MODE | sed "s/\r//")
migrations_path=$(echo $DATABASE_MIGRATIONS_PATH | sed "s/\r//")
uri="$schema://$user:$password@$host:$port/$name?sslmode=$ssl_mode"

docker compose -f docker-compose.dev.yml up database --build -d

echo -n "["
for (( i=0; i<20; i++ )); do
    echo -n "#"
    sleep 0.2
done
echo "]"
sleep 1

echo "\
+-----------------------+
| Loading migrations... |
+-----------------------+\
"
migrate -path $migrations_path -database $uri up
