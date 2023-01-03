FROM docker.io/library/golang:1.19.4-alpine3.17 as builder

COPY . /app

WORKDIR /app

RUN go build .

FROM docker.io/library/alpine:3.17.0 AS production

LABEL org.opencontainers.image.authors="Devin Canterberry <devin@canterberry.cc>"
LABEL org.opencontainers.image.url="https://github.com/twuni/pinki"
LABEL org.opencontainers.image.documentation="https://twuni.github.io/pinki"
LABEL org.opencontainers.image.source="https://github.com/twuni/pinki"
LABEL org.opencontainers.image.vendor="Twuni"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.title="Pinki"
LABEL org.opencontainers.image.description="Pinki helps developers ship software with authenticity."

COPY --from=builder /app/pinki /bin/pinki

RUN chmod 0555 /bin/pinki

ENTRYPOINT ["/bin/pinki"]
