FROM golang:1.25.1-trixie AS build
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]
