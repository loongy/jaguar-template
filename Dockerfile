FROM golang:1.8

RUN mkdir -p $GOPATH/src/github.com/loongy/jaguar
WORKDIR $GOPATH/src/github.com/loongy/jaguar
COPY . .

RUN go get ./...
RUN go install

ENV PORT "3000"
EXPOSE 3000

CMD ["jaguar"]