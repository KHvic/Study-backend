# TODO: Setup MySQL, executable will not run now because it is not setup

FROM golang:latest

LABEL maintainer="Victor Goh Ka Hian <kahiangohvictor@gmail.com>"

WORKDIR $GOPATH/src/github.com/KHvic/quiz-backend
COPY . $GOPATH/src/github.com/KHvic/quiz-backend
RUN go build .

EXPOSE 8080
ENTRYPOINT "quiz-backend.exe"
