FROM golang:1.19 AS build

WORKDIR /workdir

COPY . .

# `app/`ディレクトリが存在する場合、`-o app`を指定すると`app/${project_name}`という実行ファイルが生成される
# ディレクトリにならなさそうなapp.exeにしておく
RUN go build -ldflags "-s -w" -trimpath -mod=vendor -o app.exe .

FROM ubuntu:20.04

WORKDIR /workdir

# set timezone to Asia/Tokyo
# install ca-certificates
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y \
    tzdata \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

COPY --from=build /workdir/app.exe app.exe
CMD ["./app.exe"]
