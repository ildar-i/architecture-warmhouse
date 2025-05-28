# Load golang image to build
FROM golang:1.23-alpine as builder
ARG APP_PATH

RUN mkdir -p /app
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -buildvcs=false -o=temperature $APP_PATH


# Deploy execute file to simple linux server
FROM alpine
RUN mkdir -p /app
WORKDIR /app
COPY --chown=0:0 --from=builder /app/ ./

EXPOSE 8081

ENTRYPOINT ["/app/temperature"]