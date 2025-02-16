# FROM golang:1.20.3-bullseye
FROM golang:1.23.4-bullseye as builder
# Set the Working Directory inside the container
WORKDIR /go/src/app

# Cache and install dependencies
COPY go.mod ./
RUN go mod download

# Copy app files
COPY . .
COPY ./src/ /app/src/

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/bin/app /

EXPOSE 3000

CMD ["/app"]
