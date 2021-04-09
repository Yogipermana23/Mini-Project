# Step 1 build executable binary
FROM golang:alpine as builder

# Install git
COPY . $GOPATH/src/gitlab.bri.co.id/app/
WORKDIR $GOPATH/src/gitlab.bri.co.id/app/

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/app .

# Step 2 build a small image
# start from scratch
FROM scratch

# APM Config
##ENV ELASTIC_APM_SERVICE_NAME "BRI360-Summary"
##ENV ELASTIC_APM_SERVER_URL http://172.18.136.12:8200

# Copy our static executable
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]