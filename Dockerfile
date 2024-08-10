FROM golang:1.22 as builder

WORKDIR /flixlib-build

RUN apt-get update && \
    apt-get install -y --no-install-recommends make && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    make generate-templ && \
    make build

FROM scratch

WORKDIR /app

USER nobody:nogroup

COPY --from=builder /flixlib-build/FlixLib /app/

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD curl --fail http://localhost:8080/status || exit 1

CMD ["./FlixLib"]
