FROM golang:1.22

WORKDIR /flixlib

RUN apt-get update && \
    apt-get install -y make

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

EXPOSE 8080

CMD ["make"]
