# 编译
FROM golang:1.8.3 as builder
ADD ./ /go/src/gofamily/
RUN set -ex && cd /go/src/gofamily && CGO_ENABLED=0 go build -v -a -ldflags '-extldflags "-static"' -o gofamily 

# 打包  
FROM scratch
MAINTAINER blank <461961306@qq.com>
WORKDIR /go
COPY --from=builder /go/src/gofamily/gofamily /go/gofamily
COPY --from=builder /go/src/gofamily/conf/ /go/conf/
COPY --from=builder /go/src/gofamily/swagger/ /go/swagger/
EXPOSE 8080
ENV runmode="prod"
ENTRYPOINT ["/go/gofamily"]

