FROM golang:1.14

WORKDIR $GOPATH/src/github.com/nesitor/interview-accountapi-form3

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN chmod +x entrypoint.sh

# Run the executable
CMD ["./entrypoint.sh"]
