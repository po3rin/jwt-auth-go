# API Image 作成

# ベースとなるイメージ
FROM golang:1.10.2-alpine3.7
# FROM golang:1.10

# Gopkg.toml 等を事前にコピーして dep ensure が実行できるようにする
COPY . /go/src/auth-server/

# dep ensure を行うプロジェクトルートに移動する
WORKDIR /go/src/auth-server/api

# 必要なパッケージをイメージにインストールする
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

EXPOSE 8081
CMD go run main.go

