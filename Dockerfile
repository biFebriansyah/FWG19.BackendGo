FROM golang:1.20.3-alpine AS build

WORKDIR /goapp

COPY . .

RUN go mod download
RUN go build -v -o /goapp/goback ./cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=build /goapp /app

ENV PATH="/app:${PATH}"

EXPOSE 8081

ENTRYPOINT [ "goback" ]