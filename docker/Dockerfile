FROM alpine:latest
MAINTAINER Brian Lachniet <blachniet@gmail.com>

RUN apk add --no-cache ca-certificates
COPY . /usr/local/timetonight
WORKDIR /usr/local/timetonight
ENV PATH /usr/local/timetonight:$PATH

EXPOSE 3000
CMD ["./timetonight"]
