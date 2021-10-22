FROM golang:1.15-buster as builder

ENV GOSUMDB off
ENV GOPROXY "https://goproxy.cn,direct"
COPY . /home/http-request-print
WORKDIR /home/http-request-print
RUN go build -o http-request-print

FROM debian:buster-slim
WORKDIR /home/http-request-print
COPY --from=builder /home/http-request-print/http-request-print /home/http-request-print/
CMD ["/home/http-request-print/http-request-print"]
