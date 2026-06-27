FROM docker.io/golang:latest as build
COPY . /
WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ise-connector

## Stage 2
FROM scratch
COPY --from=build /ise-connector .
CMD [ "/ise-connector" ]
