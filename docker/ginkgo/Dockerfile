FROM golang:1.16

RUN go install github.com/onsi/ginkgo/ginkgo@v1.16.4

WORKDIR /ginkgotest

CMD ["ginkgo", "-r"]