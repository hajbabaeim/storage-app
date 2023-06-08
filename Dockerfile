# Stage 1 - Build
FROM golang:1.20.5-alpine3.18
WORKDIR /go/src/app
COPY . .  
COPY .env /go/src/app/cmd/server
WORKDIR /go/src/app/cmd/server
EXPOSE 1321
RUN go env -w GO111MODULE=auto    
RUN go env -w GOPROXY=https://goproxy.io,direct

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build

CMD ["./server"]