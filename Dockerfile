FROM golang:latest

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY binaries /app/binaries
COPY json /app/json
COPY scripts /app/scripts

RUN chmod +x /app/scripts/*

RUN apt-get update && apt-get install -y curl

EXPOSE 8080 8081

CMD ["sh", "-c", "./binaries/server & ./binaries/gateway"]