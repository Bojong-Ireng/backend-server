FROM postgres:17-alpine

COPY ./db/db_init.sql /docker-entrypoint-initdb.d/