
FROM golang:alpine
WORKDIR /go/src/app
COPY hello-fcycle.go .
RUN go build -o hello-fcycle .

CMD ./hello-fcycle