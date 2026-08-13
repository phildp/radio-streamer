FROM golang:1.26-bookworm

RUN apt-get update \
    && apt-get install -y --no-install-recommends libasound2-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Default: run the same checks as CI and `make ci`.
CMD ["make", "check"]
