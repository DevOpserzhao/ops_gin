FROM golang:1.14.2 as mod
LABEL stage=mod
ARG GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
WORKDIR /root/myapp/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM mod as builder
LABEL stage=intermediate
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gin2s3 main.go


FROM alpine:3.11.5
WORKDIR /root/app/
ARG VERSION=v0.2
COPY --from=builder /root/myapp/gin2s3 .
COPY --from=builder /root/myapp/template ./template
COPY --from=builder /root/myapp/public ./public
COPY --from=builder /root/myapp/.aws ./.aws
EXPOSE 8080
ENTRYPOINT ["./gin2s3"]