# IMAGE 1: Builder
FROM golang:1.17.5-alpine as builder

# pre-reqs
RUN apk add --no-cache git nodejs npm

# copy in src
WORKDIR $GOPATH/src/hotline/
COPY . .

# install dependencies
RUN go get -d -v

# build spa
WORKDIR $GOPATH/src/hotline/internal/app/frontend/spa
RUN npm i
RUN npm run build

# build
WORKDIR $GOPATH/src/hotline/
RUN go build -o /hotline

# IMAGE 2: Runner
FROM alpine:latest

# healthcheck dependency
RUN apk add curl

# copy binary
WORKDIR /app
COPY --from=builder /hotline .

# run
EXPOSE 8080 53 53/udp
CMD /app/hotline server