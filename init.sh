#!/bin/bash

echo "PayWay start up"

set -e

# not inside docker container
if [ ! -f "/.dockerenv" ]; then
    echo "Not running inside Docker, exiting"
    exit 0
fi

# # create .env files if not exists
# echo "Checking .env files..."
# if [ ! -f ".env" ]; then
#     echo "Creating .env from .env.example..."
#     cp .env.example .env
# else
#     echo ".env already exists, skipping..."
# fi

# if [ ! -f "app/.env" ]; then
#     echo "Creating app/.env from app/.env.example..."
#     cp app/.env.example app/.env
# else
#     echo "app/.env already exists, skipping..."
# fi

set -a
source app/.env
set +a

echo "<enviroments>"
env
echo "</enviroment>"

# DB ready
echo "Waiting for database..."
until pg_isready -h db -p "$DB_PORT" -q; do
    sleep 1
done
echo "Database is ready!"

# clean db
echo "Cleaning database..."
PGPASSWORD=$DB_PASSWORD psql -h db -U user -d payway -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;" || true

# migrations
echo "Applying migrations..."
migrate -path /project/app/$DB_MIGRATIONS_DIRECTORY -database "$DB_URL" up

echo "PayWay started"
