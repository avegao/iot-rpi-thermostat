FROM golang:1.9.0 AS build

WORKDIR /go/src/gitlab.com/avegao/iot-thermostat

COPY ./ ./

RUN go install

########################################################################################################################

FROM alpine:3.6

MAINTAINER "Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /go/bin/iot-thermostat /app/iot-thermostat

ARG VCS_REF="unknown"
ARG BUILD_DATE="unknown"

LABEL com.avegao.iot.thermostat.vcs_ref=$VCS_REF \
      com.avegao.iot.thermostat.build_date=$BUILD_DATE \
      maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENTRYPOINT ["/app/iot-thermostat"]
