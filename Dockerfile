FROM golang:1.25 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /server

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]