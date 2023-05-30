FROM golang:latest as build

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN adduser -u 10001 --disabled-password flexdb
RUN go mod download
RUN go install github.com/jteeuwen/go-bindata/go-bindata

WORKDIR /release

ADD . .

RUN go generate ./...
RUN STATIC_BUILD=1 ./build.sh
RUN CGO_ENABLED=0 go build -a ./cmd/lx

FROM alpine:3.4 as prod

COPY --from=build /release/flexdb /
COPY --from=build /release/lx /

EXPOSE 8000

USER flexdb
CMD ["/flexdb"]
