FROM golang:latest

ENV GOPROXY http://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/StephenChen/go-learning/gin/gin-example
COPY . $GOPATH/src/github.com/StephenChen/go-learning/gin/gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-example"]
