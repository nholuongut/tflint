FROM --platform=$BUILDPLATFORM golang:1.23-alpine3.20 AS builder
LABEL maintainer="Nho Luong <luongutnho@hotmail.com>"
ARG TARGETOS TARGETARCH

RUN apk add --no-cache make

WORKDIR /tflint
COPY . /tflint
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH make build

FROM alpine:3.20

LABEL maintainer=nholuongu

RUN apk add --no-cache ca-certificates

COPY --from=builder /tflint/dist/tflint /usr/local/bin

ENTRYPOINT ["tflint"]
WORKDIR /data
