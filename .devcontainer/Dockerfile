FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu-20.04

RUN set -x \
    && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo 'Asia/Tokyo' >/etc/timezone

# golang
ARG GO_VERSION=1.19
RUN set -x && curl -fsSL "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C /usr/local
ENV PATH $PATH:/usr/local/go/bin

USER vscode
RUN go env -w GOPRIVATE=github.com/kurusugawa-computer

# RUN sudo apt-get update && sudo apt-get install -y \
#     [something to install]
#  && sudo apt-get clean \
#  && sudo rm -rf /var/lib/apt/lists/*
