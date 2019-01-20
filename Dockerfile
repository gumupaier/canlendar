FROM golang:latest

MAINTAINER henson_wu "henson_wu@foxmail.com"

WORKDIR $GOPATH/src/calendar
ADD . $GOPATH/src/calendar

RUN go get github.com/astaxie/beego

RUN go build .

EXPOSE 8080

ENTRYPOINT ["/work/script/init.sh"]
ENTRYPOINT ["./calendar"]
