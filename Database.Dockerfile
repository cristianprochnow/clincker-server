FROM mysql:5.7
EXPOSE 3306
COPY ./db/dump/ /docker-entrypoint-initdb.d/
