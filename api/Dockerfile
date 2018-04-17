FROM golang:latest AS build

WORKDIR /go/src/github.com/flexphere/sts

RUN set -e \
	&& rm -rf vendor \
	&& go get -u github.com/golang/dep/cmd/dep

COPY . .

RUN set -e \
	&& dep ensure \
	&&  CGO_ENABLED=0 go build -ldflags="-s -w" .


FROM debian:latest

COPY --from=build /go/src/github.com/flexphere/sts/sts /bin/sts

CMD ["/bin/sts"]