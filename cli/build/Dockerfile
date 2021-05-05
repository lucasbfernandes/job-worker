FROM golang:1.16-alpine

RUN apk --no-cache add gcc musl-dev git mercurial

WORKDIR /app/cli

COPY . .

# Running lint
RUN (wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.39.0 && \
     cp ./bin/golangci-lint /bin/ && \
     rm -r ./bin)

RUN golangci-lint run

# Building app
RUN go build -o ./out/cli cmd/cli/*.go

CMD ["./out/cli"]
