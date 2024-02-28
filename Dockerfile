FROM golang:1.22

WORKDIR /flixlib

RUN apt-get update && \
    apt-get install -y ca-certificates make curl unzip && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN go install github.com/cosmtrek/air@latest && \
    curl -fsSL https://bun.sh/install | bash && \
    go mod download

ENV BUN_INSTALL="/root/.bun"
ENV PATH="$BUN_INSTALL/bin:$PATH"

EXPOSE 443

CMD ["make"]
