FROM golang:1.23-alpine

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o musicApp ./cmd


FROM alpine:latest

ENV DB_HOST=
ENV DB_PORT=
ENV DB_USERNAME=
ENV DB_PASSWORD= 
ENV DB_SSL= 
ENV LOG_LVL= 
ENV API_URL= 

COPY --from=0 /app/musicApp /musicApp
COPY ./build/entrypoint.sh /entrypoint.sh
COPY ./scripts/migrations /migrations
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]