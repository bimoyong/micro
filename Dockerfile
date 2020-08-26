FROM --platform=$TARGETPLATFORM alpine:latest
RUN apk --no-cache --virtual .build_deps add make git go gcc libtool musl-dev

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . /
RUN make

# target multiple architecture
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} make

RUN apk add ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/* && \
    [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

# cleanup to optimize image size
RUN go clean --i --n --r --x --cache --testcache --modcache
RUN apk del --purge .build_deps
RUN rm -rf /go

ENTRYPOINT ["/micro"]
