FROM golang:1.19
WORKDIR /qqchatgpt
COPY . .
RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install gcc g++ ca-certificates chromium xvfb && \
    go mod download && \
    go build qqchatgpt.go
ENTRYPOINT [ "/qqchatgpt/qqchatgpt" ]