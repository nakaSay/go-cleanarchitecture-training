FROM golang:1.17.0-bullseye

WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPATH /go

RUN apt-get update \
    && apt-get install -y git \
    && go get -v \
        github.com/urfave/cli/v2 \
        github.com/go-delve/delve/cmd/dlv@latest \
        github.com/rogpeppe/godef@latest \
        github.com/ramya-rao-a/go-outline@latest \
        github.com/labstack/echo/v4 \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

RUN go install golang.org/x/tools/gopls@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN GO111MODULE=off go get github.com/oxequa/realize