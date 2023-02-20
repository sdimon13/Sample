FROM golang:latest AS build

# install migrate to perform DB migration
WORKDIR /app/backend
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz && mv migrate /usr/bin/

# install go-swagger
RUN curl -o /usr/local/bin/swagger -L'#' "https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64" \
    && chmod +x /usr/local/bin/swagger

# install protoc
RUN apt update && apt install -y protobuf-compiler
#RUN curl -o /usr/local/bin/protoc -L'#' "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip" \
#    && chmod +x /usr/local/bin/protoc

RUN go clean --cache

# copy module files first so that they don't need to be downloaded again if no change
COPY go.* ./
RUN go mod download
RUN go mod verify

RUN go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

# copy source files and build the binary
COPY . .

RUN make proto

RUN make build
CMD ["/sample", "api"]