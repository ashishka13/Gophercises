FROM golang

WORKDIR /go/src/myapp
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["./main"]