FROM golang:1.11-stretch as build

WORKDIR /go/src/github.com/vmj/fixed-response

COPY fixed-response.go ./
RUN CGO_ENABLED=0 go build -a -o fixed-response

FROM scratch
COPY --from=build /go/src/github.com/vmj/fixed-response/fixed-response /
CMD ["/fixed-response"]
