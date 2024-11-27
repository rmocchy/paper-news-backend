FROM node:18-alpine

# 必要なツールのインストール
RUN npm install -g @apidevtools/swagger-cli

# 作業ディレクトリを設定
WORKDIR /app

# デフォルトのエントリーポイントを設定
ENTRYPOINT ["swagger-cli"]
