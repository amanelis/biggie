FROM golang:latest

#ENV SKYNET__ENVIRONMENT local
#ENV SKYNET__COINBASE_KEY {YOUR_KEY}
#ENV SKYNET__COINBASE_SECRET {YOUR_SECRET}
#ENV SKYNET__COINBASE_PHRASE {YOUR_PHRASE}

ADD . /go/src/github.com/amanelis/skynet

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go install github.com/amanelis/skynet

ENTRYPOINT /go/bin/skynet
