FROM golang:1.18-alpine as build

RUN mkdir /src
WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . .

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN CGO_ENABLED=${CGO_ENABLED} \
    GOOS=${GOOS} \
    GOARCH=${GOARCH} \
    go build -o simple-api -ldflags "-w -s" main.go


FROM scratch
WORKDIR /src
COPY --from=build /src/simple-api /src/simple-api

EXPOSE 8000

ENTRYPOINT ["/src/simple-api"]
