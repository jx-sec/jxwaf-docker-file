FROM golang:1.18 as builder

WORKDIR /opt/app
COPY ./go.mod .
COPY ./go.sum .
# RUN mkdir -p /usr/local/go/src/
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN  go build  -o /opt/app/main ./main.go

FROM gcr.io/distroless/base-debian11
# from alpine:3.16
WORKDIR /opt/app/
COPY --from=builder /opt/app/main .
COPY --from=builder /opt/app/GeoLite2-City.mmdb .
# run   ls /opt/app/main -al
# kafaka兼容要求
# run mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# EXPOSE 30882

ENTRYPOINT  ["/opt/app/main"]
