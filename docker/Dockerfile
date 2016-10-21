FROM alpine:latest
MAINTAINER Brian Lachniet <blachniet@gmail.com>

# Install root CA certificates
RUN apk add --no-cache ca-certificates

# Put the timezone database in the standard Go location
COPY zoneinfo.zip /usr/local/go/lib/time/

# Copy applicaiton files
COPY templates /usr/local/timetonight/templates
COPY timetonight /usr/local/timetonight/
WORKDIR /usr/local/timetonight
ENV PATH /usr/local/timetonight:$PATH

EXPOSE 3000
CMD ["timetonight"]
