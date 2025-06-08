FROM golang:1.24 as build

WORKDIR /app

COPY go.mod go.sum ./

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg

# building go with external linking for support of cgo
# building statically linked executable
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'"  -o main ./cmd/app/main.go

FROM scratch AS release

WORKDIR /app

COPY --from=build /app/main /main

CMD [ "/main" ]
