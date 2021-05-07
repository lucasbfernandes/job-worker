# job-worker

Simple Job Worker service that provides an API to run arbitrary Linux processes.

## Setup
In order to run the CLI and Server inside a docker environment, please install [docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/).
After it, from the job-worker's root directory, execute the following command:

```
docker-compose build && docker-compose up
```

###  Server

<strong>Running the application:</strong>

```
docker-compose exec -T server ./out/server
```

(Will be available at http://localhost:8080 and a Postman collection can be found [here](assets/postman)):


<strong>Running tests:</strong>

```
docker-compose exec -T server go test -v ./...
```

<strong>Running everything outside a docker environment:</strong>

You can also run and test the server outside a docker environment. In order to do that, [install go 1.16](https://golang.org/doc/install) and execute the following commands from the server's root directory:

1. `go test -v ./...`
2. `go build -o ./out/server cmd/server/*.go`
3. `LOGS_DIR=<path-to-logs-dir> ./out/server -port=8080`

<strong>PS:</strong>

The Server uses the `LOGS_DIR` environment variable to create the folder that will contain process logs files. This is the
absolute path of the directory, and it won't be created again if it already exists. It is important to note that the server
must have enough permissions to create the dir in that location.

Example:

`LOGS_DIR = /app/server/logs`

Resulting files:

* /app/servers/logs/\<job-id\>

If LOGS_DIR is not provided, files will be created inside the folder `logs` in the current directory.

### CLI

<strong>Running the application:</strong>

Before executing the CLI, please have the server up and running.

```
docker-compose exec -T server ./out/server
```

In another terminal window, execute the next command to hop in the CLI container shell:
```
docker-compose exec -T cli /bin/sh
```

Execute one of the following commands:
```
Create Job:
./out/cli exec -s SERVER_URL -t API_TOKEN EXECUTABLE [ARGS...]

List Jobs:
./out/cli list -s SERVER_URL -t API_TOKEN

Stop Job:
./out/cli stop -s SERVER_URL -t API_TOKEN -i JOB_ID

Get Job Status:
./out/cli status -s SERVER_URL -t API_TOKEN -i JOB_ID

Get Job Logs:
./out/cli logs -s SERVER_URL -t API_TOKEN -i JOB_ID
```

<strong>PS:</strong>

SERVER_URL will default to http://server:8080 but you can override it with the `-s` flag.

The Job Worker has 2 available users. One with the `ADMIN` role and another with the `USER` role.
Both have access to every resource, with the only difference being that the `ADMIN` can interact with all resources,
even those created by other users.

These are the tokens associated with the users:

* ADMIN - `qTMaYIfw8q3esZ6Dv2rQ`
* USER - `9EzGJOTcMHFMXphfvAuM`

Please provide them to the `-t` flag.

<strong>Running tests:</strong>

```
docker-compose exec -T cli go test -v ./...
```

<strong>PS:</strong> All CLI tests need a mock server up and running. This is the stubby4j container started with `docker-compose up`. Running tests inside the docker container will make
DEFAULT_SERVER_MOCK_URL be evaluated as `http://stubby4j:8883`.

<strong>Running everything outside a docker environment:</strong>

You can also run and test the CLI outside a docker environment. In order to do that, [install go 1.16](https://golang.org/doc/install) and execute the following commands from the CLI's root directory:

1. `go test -v ./...`
2. `go build -o ./out/cli cmd/cli/*.go`
3. `./out/cli <command>`

<strong>PS:</strong>

The CLI needs a server running, and you can provide its url using the `-s` flag in each command or by putting it inside the `DEFAULT_SERVER_URL` environment variable.

The `stubby4j` container must be running for the tests, therefore you still need to run `docker-compose up stubby4j` before it.
