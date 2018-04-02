FROM golang:1.7.3
WORKDIR /go/src//go/src/github.com/splusminusx/tcp-probe
COPY tcp_probe.go .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tcp_probe tcp_probe.go

FROM cloudprober/cloudprober:latest
COPY --from=0 /go/src//go/src/github.com/splusminusx/tcp-probe/tcp_probe /probes/tcp_probe
