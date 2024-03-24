FROM golang:1.21

WORKDIR /go/delivery

RUN apt update
RUN apt install unzip

COPY . .
RUN make build

# install psql
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o flood-control-task ./cmd/main.go

CMD ["./bin/app", "host.docker.internal:8086", "-app_port", "8082"]

