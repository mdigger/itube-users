#!/usr/bin/env sh

FNAME=./_migrate

if [ ! -f $FNAME ]; then
    go get -u -d github.com/golang-migrate/migrate/cli github.com/lib/pq
    go build -tags 'postgres' -o $FNAME github.com/golang-migrate/migrate/cli
fi;

$FNAME $@