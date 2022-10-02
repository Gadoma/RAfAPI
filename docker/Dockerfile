FROM --platform=linux/amd64 alpine:latest

LABEL maintainer="Piotr Gadziński <dev@gadoma.com>" 

WORKDIR /app

COPY bin/rafapi /app/rafapi

EXPOSE 5000

ENTRYPOINT [ "/app/rafapi" ]
