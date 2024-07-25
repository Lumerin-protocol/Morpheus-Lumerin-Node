FROM golang:1.22.3 as base
RUN apt-get update && apt-get install -y procps curl gnupg2 dbus-x11

FROM base as builder
ARG COMMIT
ENV COMMIT=$COMMIT

WORKDIR /app 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./build.sh

FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/proxy-router /usr/bin/
COPY --from=builder /usr/bin/dbus-launch /usr/bin/
COPY --from=builder /app/.env /app/.env
COPY --from=busybox /bin /bin
COPY --from=busybox /lib /lib

SHELL ["/bin/sh", "-c"]
EXPOSE 3333 8080 8082
ENTRYPOINT ["proxy-router"]