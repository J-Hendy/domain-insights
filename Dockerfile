FROM golang:1.13.3-alpine3.10 AS build

RUN apk add --no-cache git openssh

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app/domain-insights/

COPY . .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

RUN go build -o bin/domain-insights .

FROM alpine:3.10

RUN apk add --no-cache bash ca-certificates curl

RUN adduser -D -u 1000 domusr

WORKDIR /home/domusr

USER domusr

COPY --from=build /app/domain-insights/bin/domain-insights/ .

EXPOSE 8080

CMD ./domain-insights