FROM golang:1.16-alpine

RUN apk --no-cache add gcc musl-dev git mercurial

WORKDIR /app/server

COPY . .

# Running lint
RUN (wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.39.0 && \
     cp ./bin/golangci-lint /bin/ && \
     rm -r ./bin)

RUN golangci-lint run

# Building app
RUN go build -o ./out/server cmd/server/*.go

EXPOSE 8080

ENTRYPOINT ["./out/server"]