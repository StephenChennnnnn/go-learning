FROM golang:latest

ENV GOPROXY http://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/StephenChen/go-learning
COPY . $GOPATH/src/github.com/StephenChen/go-learning/gin/gin-example
RUN go mod init
RUN go build ./gin/gin-example

EXPOSE 8000
ENTRYPOINT ["./gin-example"]
