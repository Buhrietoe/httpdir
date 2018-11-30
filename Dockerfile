FROM golang:1-alpine AS build
COPY . /go/src/github.com/Buhrietoe/httpdir/
WORKDIR /go/src/github.com/Buhrietoe/httpdir/
ENV CGO_ENABLED 0
RUN go version && \
    go build -v -ldflags "-s -w" -o httpdir .

FROM scratch
LABEL maintainer "Jason Gardner <buhrietoe@gmail.com>"
EXPOSE 8080
COPY --from=build /go/src/github.com/Buhrietoe/httpdir/httpdir /httpdir
ENTRYPOINT ["/httpdir"]
