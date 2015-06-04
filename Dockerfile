FROM gliderlabs/alpine:3.1
ENTRYPOINT ["/bin/irukad"]

COPY . /go/src/github.com/spesnova/iruka
RUN apk-install -t build-deps go git mercurial \
      && cd /go/src/github.com/spesnova/iruka/irukad \
      && export GOPATH=/go \
      && go get \
      && go build -ldflags "-X main.Version" -o /bin/irukad \
      && rm -rf /go \
      && apk del --purge build-deps go git mercurial
