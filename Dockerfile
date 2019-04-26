FROM golang:alpine
ENV APP_NAME user-server

WORKDIR $GOPATH/src/$APP_NAME
COPY . $GOPATH/src/$APP_NAME

RUN apk add --update --no-cache make \
    && cd $GOPATH/src/$APP_NAME \
    && make

WORKDIR /app
VOLUME /app/config

RUN mv $GOPATH/src/$APP_NAME/release/* ./ \
    && rm -rf $GOPATH/src/$APP_NAME

CMD ["/app/user-server"]
