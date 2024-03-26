FROM ubuntu:latest
COPY go-common-libs-test .
ENTRYPOINT ./go-common-libs-test
