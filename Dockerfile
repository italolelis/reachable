FROM golang:1.11beta3-alpine AS builder

WORKDIR ./reachable
COPY . ./

RUN apk add --update alpine-sdk
RUN make all

# ---
FROM alpine

COPY --from=builder ./go/reachable/dist/reachable ./

ENTRYPOINT ["/reachable"]
