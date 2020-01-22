FROM golang

ENV APP_NAME myproject
ENV PORT 8080

WORKDIR /go/src/github.com/thezultimate/${APP_NAME}/
COPY . .

RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}