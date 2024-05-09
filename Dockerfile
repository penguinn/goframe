# base image
FROM iregistry.baidu-int.com/acg-det/golang:1.21.6

COPY output/ /go/bin/
WORKDIR /go/bin/
CMD ["goframe"]