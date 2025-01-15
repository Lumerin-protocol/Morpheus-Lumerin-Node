# Stage 1: Build
FROM golang:1.22.3-alpine AS builder

# Capture the Git tag and commit hash during build
ARG TAG_NAME
ARG COMMIT
ENV TAG_NAME=$TAG_NAME
ENV COMMIT=$COMMIT

WORKDIR /app
COPY . .

# Build the Go binary (Docker will build automatically to the OS and Arch that is hosting)
RUN CGO_ENABLED=0 \
  TAG_NAME=$TAG_NAME COMMIT=$COMMIT ./build.sh && \
  cp /bin/sh /app/sh && chmod +x /app/sh

# Stage 2: Final Image
FROM scratch
WORKDIR /app

# Copy required files and binary
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/proxy-router /usr/bin/
COPY models-config.json /app/data/models-config.json

# Optional Copy utilities from busybox image
COPY --from=busybox /bin /bin
COPY --from=busybox /lib /lib

SHELL ["/bin/sh", "-c"]
EXPOSE 3333 8082

ENTRYPOINT ["proxy-router"]
