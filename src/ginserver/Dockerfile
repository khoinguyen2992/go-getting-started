FROM khailq/mybb-go-base:0.1-alpine
MAINTAINER Khoi Nguyen <minhkhoi@siliconstraits.com>

WORKDIR /go/src/ginserver

RUN go get github.com/tools/godep
COPY ./Godeps /go/src/ginserver/Godeps
RUN godep restore

COPY . /go/src/ginserver
RUN chmod 775 /go/src/ginserver/docker-start.sh

CMD sh ./docker-start.sh