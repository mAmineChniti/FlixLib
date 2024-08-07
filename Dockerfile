FROM golang:1.22

WORKDIR /flixlib

RUN apt-get update && \
    apt-get install -y ca-certificates make curl unzip && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN go install github.com/air-verse/air@latest && \
    go mod download

EXPOSE 443

CMD ["make"]
