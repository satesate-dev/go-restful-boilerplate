FROM alpine:3.12.0
WORKDIR /root

RUN set -ex && \
    apk add --no-cache gcc musl-dev
RUN mkdir file
RUN mkdir log

COPY server server
CMD ./server --config=.config.toml >> log/stdout.log 2>> log/stderr.log
