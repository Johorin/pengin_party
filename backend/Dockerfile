FROM golang:1.25.4

WORKDIR /go/pengin_party

# wireのインストール
RUN go install github.com/google/wire/cmd/wire@latest

# ホットリロードツールのインストール
RUN go install github.com/air-verse/air@latest

# CLI migrationツールのインストール @see: https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENV GIN_MODE-debug
ENV APP_ENV=local
ENV TZ="Asia/Tokyo"

# 対話セッションを維持するためにシェルを起動
CMD ["sh"]