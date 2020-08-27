FROM --platform=$TARGETPLATFORM golang:1.14.7-alpine AS buildstage
RUN apk add --no-cache make git

COPY . /

# target multiple architecture
ARG TARGETOS
ARG TARGETARCH

RUN cd / && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} make

FROM --platform=$TARGETPLATFORM alpine:latest AS mainstage
RUN apk add ca-certificates && \
    [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

COPY --from=buildstage /micro /

ENTRYPOINT ["/micro"]
