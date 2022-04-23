FROM golang:1.18 as base

FROM base as dev

WORKDIR /usr/src/app

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]

FROM base AS build

WORKDIR /usr/src/app

COPY . .

RUN go build -o /runner

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /runner /runner

USER nonroot:nonroot

ENTRYPOINT ["/runner"]



