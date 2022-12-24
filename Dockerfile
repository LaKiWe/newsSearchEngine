FROM golang:1.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io

COPY . /app
WORKDIR /app

RUN go get && go build -ldflags="-s -w" -installsuffix cgo

FROM debian:buster-slim

ENV TZ=Asia/Shanghai \
    LANG=C.UTF-8 \
    APP_DIR=/usr/local/news_SearchEngine

COPY --from=builder /app/newsSearchEngine ${APP_DIR}/newsSearchEngine

WORKDIR ${APP_DIR}

RUN ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && chmod +x gofound

EXPOSE 5678

CMD ["./newsSearchEngine"]