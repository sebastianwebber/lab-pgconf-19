FROM golang:1.11 AS build

    ARG GIT_TAG
    ARG GIT_COMMIT

    ENV WORKDIR=/build
    ENV GOOS=linux
    ENV CGO_ENABLED=0
    ENV GOARCH=amd64

    WORKDIR ${WORKDIR}
    COPY . /build
    RUN mkdir dist && \
        go build -o dist/desafio cmd/main.go && \
        chmod +x dist/*


FROM drone/ca-certs
    ENV GIN_MODE=release
    
    EXPOSE 8080

    WORKDIR /app

    COPY --from=build /build/dist/desafio .

    CMD [ "/app/desafio" ]