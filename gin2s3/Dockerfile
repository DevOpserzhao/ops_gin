FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/DevOpserzhao/ops_gin/gin2s3
COPY . $GOPATH/src/github.com/DevOpserzhao/ops_gin/gin2s3
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./gin2s3"]