# https://github.com/GoogleContainerTools/distroless#examples-with-docker

# https://hub.docker.com/_/golang
FROM golang:1.19-alpine AS build
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY app/ app/
COPY main.go .
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -o /go/bin/app

# https://github.com/GoogleContainerTools/distroless/tree/main/base
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
# COPY --from=build /go/src/app/app/sa.json / # for local
CMD ["/app"]
