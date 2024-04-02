FROM mysql:5.7

COPY ./db/dump/ /docker-entrypoint-initdb.d/
