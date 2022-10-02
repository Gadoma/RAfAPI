FROM --platform=linux/amd64 alpine:latest

LABEL maintainer="Piotr Gadziński <dev@gadoma.com>" 

WORKDIR /app

COPY bin/rafapi /app/rafapi
COPY db/db.dist.sqlite /app/db/db.dist.sqlite

EXPOSE 5000

ENTRYPOINT [ "/app/rafapi" ]
