# ベースイメージを指定
FROM golang:alpine

# 作業ディレクトリを設定
WORKDIR /app

# ホストOS上のソースコードをコンテナ内にコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
