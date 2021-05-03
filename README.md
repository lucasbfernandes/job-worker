# job-worker

Simple Job Worker service that provides an API to run arbitrary Linux processes.

###  Server
In order to run the Server, please install [docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/).

From the root directory, execute `docker-compose build && docker-compose up` and choose one of the following commands in another terminal window:

* Run the application (Will be available at http://localhost:8080 and a Postman collection can be found [here](assets/postman)):

        docker-compose exec -T server ./out/server

* Run tests:

        docker-compose exec -T server go test -v ./...

You can also run and test the application outside a docker environment. In order to do that, [install go 1.16](https://golang.org/doc/install) and execute the following commands from the root directory:

1. `go test -v ./...`
2. `go build -o ./out/server cmd/server/*.go`
3. `LOGS_DIR=<path-to-logs-dir> ./out/server`

PS: The Server uses the `LOGS_DIR` environment variable to create the folder that will contain process logs files.