FROM golang:latest AS build

WORKDIR /app

# install migrate to perform DB migration TODO: move to base image
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz && mv migrate /usr/bin/

# install go-swagger TODO: move to base image
RUN curl -o /usr/local/bin/swagger -L'#' "https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64" \
    && chmod +x /usr/local/bin/swagger

# copy module files first so that they don't need to be downloaded again if no change
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY .. .

RUN make swagger

RUN make build

FROM scratch as final
WORKDIR /app

COPY --from=build /sample /app/sample
COPY --from=build /app/third_party /app/third_party
COPY --from=build /app/api/api.swagger.json /app/api/api.swagger.json

CMD ["/app/sample"]