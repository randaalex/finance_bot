##
## Build
##
FROM golang:1.16-alpine as build

LABEL maintainer="Alex Randa"

ADD . /app
WORKDIR /app

#RUN apk add build-base
RUN apk add --no-cache make

ENV GOPRIVATE github.com/randaalex

RUN make build

##
## Deploy
##
FROM golang:1.16-alpine AS base

WORKDIR /app

COPY --from=build /app/app.env.tmpl ./app.env
COPY --from=build /app/bin ./bin

CMD ["./bin/finance_bot", "mail"]
