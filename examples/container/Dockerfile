FROM golang:1.18 AS golang

COPY ./cmd /go/src/opitz-consulting.com/examples/container/cmd
COPY ./internal /go/src/opitz-consulting.com/examples/container/internal
COPY ./go.mod  /go/src/opitz-consulting.com/examples/container/go.mod
COPY ./go.sum  /go/src/opitz-consulting.com/examples/container/go.sum

WORKDIR /go/src/opitz-consulting.com/examples/container

RUN go install ./...
RUN CGO_ENABLED=0 go build -o my_service cmd/container/main.go

FROM busybox

COPY --from=golang /go/src/opitz-consulting.com/examples/container/my_service /app/my_service

EXPOSE 8081
CMD ["/app/my_service"]