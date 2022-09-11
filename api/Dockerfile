FROM golang:1.14.9

WORKDIR /go/src/
RUN mkdir bin
COPY ./bin/main ./bin
# RUN CGO_ENABLES=0 GOOS=linux go build -a -installsuffix cgo -o bin/main
CMD ["./bin/main"]