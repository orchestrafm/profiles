# Set Builder Image
FROM golang:1.12.17-alpine3.10 as builder

# Add Build Dependencies and Working Directory
RUN apk --no-cache add build-base git
RUN GO111MODULE=on go install github.com/gobuffalo/packr@v1.30.1
RUN GO111MODULE=on go install github.com/go-task/task@v2.0.0
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Compile
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -i -v -a -installsuffix cgo -ldflags '-extldflags "-static"' -o service ./src/

# Move to Base Image and Run
FROM alpine:3.12.0
RUN apk update && apk upgrade && apk add ca-certificates
COPY --from=builder /build/service /app/
WORKDIR /app
CMD ["./service"]