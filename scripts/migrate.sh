#!/usr/bin/env bash

# Load .env into the environment
if [ -f .env ]; then
    set -a
    source .env
    set +a
fi

command=$1
name=$2

case "$command" in
    "up")
        migrate -path migrations -database "$DATABASE_URL" up
        ;;
    "down")
        count=${name:-1}
        echo "Rolling back $count migration(s). Continue? [y/N]"
        read -r confirm
        if [ "$confirm" = "y" ]; then
            migrate -path migrations -database "$DATABASE_URL" down "$count"
        fi
        ;;
    "create")
        migrate create -ext sql -dir migrations -seq "$name"
        ;;
esac
