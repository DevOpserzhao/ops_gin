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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o k8s-deploy main.go


FROM alpine:3.11.5
WORKDIR /root/app/
ARG VERSION=v0.2
COPY --from=builder /root/myapp/k8s-deploy .

EXPOSE 8000
ENTRYPOINT ["./k8s-deploy"]