# Job Worker Design Document

This document describes the key design choices and trade-offs for the development of the Job Worker, an application
that enables the execution of arbitrary linux processes remotely. <strong>Users that interact with the system are able to start, stop,
query status and also read the stdout logs</strong> of any requested process execution.

## Architecture overview

The application is composed of 2 components: The CLI and the Server. The former is responsible for parsing user input and translating it
into REST HTTPS requests. The Server receives such requests, maintains internal in-memory state and manages the pool of linux processes that are spawned
along with their status.

It's possible to see numbers 1, 2 and 3 in the image below, each representing a different step of a user request. Each step of the request will be described in
details later in this document, but for now, here is a summary of what's happening:

* <strong>Step 1</strong>: The user types a CLI command requesting one of the possible
  interactions with a linux process. An API Token will be provided as a CLI flag.

* <strong>Step 2</strong>: The CLI application parses the user's command and translates it into an HTTPS request for the Server. The API module of the Server receives
  the request, checks if the input is valid and verifies if the user is authorized to perform the requested operation.

* <strong>Step 3</strong>: If all validations pass, the API module will forward the request to the corresponding internal handler [1].
  It will do business logic, access the persistence layer and invoke the OS Process API with the Worker Library when necessary.

![Architecture](../assets/images/architecture.png)

<strong>PS:</strong> The image depicts a situation where the CLI and the Server are running on separate machines/OS, which is not always the case. This was intended to highlight the necessity
of creating a secure channel between them.

The next sections will describe in details each of key aspects of the 2 components of the system. The design decisions and trade-offs will be shown in a way that by the end of the
document the reader will have a "grey-box" understanding of the entire system (i.e. Some internal nuances will be mentioned, but not all of them for the sake of readability. These
will be further discussed during application development.)

## CLI

The Job Worker CLI is the user interface of the Job Worker Service. Its responsibility is to authenticate the user with the Server, parse commands, generate authenticated HTTPS requests
and exhibit the responses in a structured manner. This section will:

1. Describe which commands are available for the user and how to use them;
2. Describe how the CLI will manage user secrets;
4. Show the trade-offs and what could and should be done for future work.

### Commands

* [Create Job](cli/jobs/create-job.md): `job-worker exec -s SERVER_URL -t API_TOKEN EXECUTABLE [ARGS...]`
* [List Jobs](cli/jobs/list-jobs.md): `job-worker list -s SERVER_URL -t API_TOKEN`
* [Stop Job](cli/jobs/stop-job.md): `job-worker stop -s SERVER_URL -t API_TOKEN -i JOB_ID`
* [Get Job Status](cli/jobs/get-status.md): `job-worker status -s SERVER_URL -t API_TOKEN -i JOB_ID`
* [Get Job Logs](cli/jobs/get-logs.md): `job-worker logs -s SERVER_URL -t API_TOKEN -i JOB_ID`

### Managing User Secrets

For the sake of simplicity, the CLI will receive an api token input each time a command is invoked. This api token will be matched against a set of seed api tokens,
and if a match happens, a hardcoded JWT token will be sent in the form `Authorization: Bearer <jwtToken>`. This token will be used to authenticate requests with the Server and
its generation will be explained in Server Security section.

### Trade-offs

* Few command options;
* Password won't be required since we are mocking JWT tokens.

### Future work

* Implement more CLI commands options;
* Implement login step.

## Server

The Job Worker Server is responsible for receiving HTTPS requests, applying validations, and executing the requested action if possible. This section will:

1. Describe which actions are available for requests and how the Server expects to receive them;
2. Describe how the Server will handle security concerns (Authentication and Authorization);
3. Explain briefly how the Server will keep track/state of every spawned process;
4. Show the trade-offs and what could and should be done for future work.

### REST API

#### Jobs:
* [Create Job](api/jobs/create-job.md): `POST /jobs`
* [Get Jobs](api/jobs/get-jobs.md): `GET /jobs`
* [Stop Job](api/jobs/stop-job.md): `POST /jobs/:id/stop`
* [Get Job Status](api/jobs/get-status.md): `GET /jobs/:id/status`
* [Get Job Logs](api/jobs/get-logs.md): `GET /jobs/:id/logs`

### Security

The Job Worker Service will use HTTPS + Bearer Authentication (JWT) in its initial version. Authorization
will take form with the following roles: `ADMIN` and `USER`.

#### Transport Layer Security

All communication between CLI and Server must happen in a secure channel, and for that,
we will use HTTPS with TLS 1.3. It is the most recent protocol release and has several improvements
when comparing with 1.2, including:

* Some algorithms and ciphers that are theoretically and practically vulnerable were removed (e.g. RC4 Stream Cipher, 3DES);
* Faster handshake (1 less RTT);
* Eliminates RSA key exchange;
* etc.

HTTPS/TLS will be configured in the Server with the following steps:

1. Both keys will be generated:
    * Private key: `openssl genrsa -out server.key 2048`
    * <strong>Self-signed</strong> public key: `openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`
  
2. Both keys will be stored inside the git repository (This is not a good practice but will happen because of time constraints).
3. Finally, Golang's `http.listenAndServeTLS` function will start the HTTPS server <strong>(PS: It defaults to  TLS version 1.3)</strong>:
   ```
      err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
      if err != nil {
          log.Fatal("Failed to start server: ", err)
      }
   ```

#### Authentication

Every request made to the Server must contain an authorization header in the form `Authorization: Bearer <jwtToken>`,
where `<jwtToken>` is a JWT token with an `apiToken` claim. The token will be signed using the RSA algorithm
and SHA-512 hash algorithm.

Claim example:
```
{
  "apiToken": "6q6Tz5NBELFo5E9iOSEo"
}
```

After receiving the request, the Server will check if the token is valid and if the `apiToken` claim matches any user's api token in
its in-memory database. If the request is not valid, a `401 Unauthorized` will be returned.

As explained before, JWT tokens will be mocked in the initial Job Worker version. Since there won't be a login step, tokens
will be mocked inside the CLI and will be generated with the following steps:

1. A pair of public/private keys will be generated:
    * Private key: `openssl genrsa -out server.key 2048`
    * <strong>Self-signed</strong> public key: `openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`
2. Private key will be used to create RS512 JWT tokens on https://jwt.io/.
3. Public key will be stored inside the git repository and will be used by the [jwt-go](https://github.com/dgrijalva/jwt-go) lib 
to verify incoming tokens on the server (Private key won't be stored inside the git repository).

<strong>PS:</strong>

* JWT key pairs will be different from the ones used for TLS;
* Passwords won't be stored in the initial version. A login step with username/password must be created in future versions.

#### Authorization

Authorization will be handled with a simple RBAC mechanism. There will be 2 types of roles, each with
a set of permissions associated with it.

<strong>PS: Stub users will be created as seed data when the Server starts. There will be no users CRUD in the initial version, and for that reason there won't be any roles for user management.</strong>

|  Roles               | Description |
| :-------------------:| :-----------: |
|  ADMIN | Can access every API resource and can interact with jobs created by other users|
|  USER | Can access every API resource but can only interact with resources he/she created |

When receiving an authenticated request, the Server will always check if the user has enough permissions to access the resource. If not, `403 Forbidden` will be returned.

<strong>PS: For now, role/user relationship will be one to many.</strong>

### Managing Linux Processes

The Golang exec package [5] will be responsible for the interaction with the operating system. This package will be used by the Server in a library that could be imported by
any Golang project that  wishes to do any of the following actions on linux processes: `create`, `stop`, `read logs`.

In order to keep track of processes (jobs) data, the Server will persist their data and update it according to the outputs of the `exec` package commands.

Jobs will have 4 possible states:

* <strong>RUNNING:</strong> Process is currently running with no errors;
* <strong>FAILED:</strong> Process finished with errors;
* <strong>STOPPED:</strong> Process was forced to stop by a user;
* <strong>COMPLETED:</strong> Process finished without errors.

Job states will be stored in the job object inside the in-memory database. One goroutine will be created for each job, and each will perform a database write/update on a
different record.

<strong>Stopping jobs:</strong>

For simplicity, when a user requests a job to stop, <strong>the Server will always send a `SIGKILL` signal to it right away</strong>. A future improvement would be sending a `SIGTERM` signal in an attempt to gracefully stop it without creating any zombie processes.

```
  // Start a process:
  cmd := exec.Command("sleep", "5")
  if err := cmd.Start(); err != nil {
      log.Fatal(err)
  }

  // Killing it with SIGKILL:
  if err := cmd.Process.Kill(); err != nil {
      log.Fatal("failed to kill process with SIGKILL: ", err)
  }
  
```

<strong>PS: This should be interpreted as pseudocode. It is not meant to represent the actual code.</strong>

<strong>Saving jobs logs:</strong>

The Job Server will persist almost everything inside its in-memory database, except logs. These will be saved in files inside the 
Server's executable folder in order to optimize primary memory usage. Logs will be fetched and saved in the following way:

```
cmd := exec.Command("sh", "-c", "echo hello")
stdoutPipe, _ := cmd.StdoutPipe()
stderrPipe, _ := cmd.StderrPipe()
if err := cmd.Start(); err != nil {
   // handle error
}

jobStdoutLogFile, err := os.Create("<job-id>-out")
// handle err
defer jobStdoutLogFile.Close()
_, err = io.Copy(jobStdoutLogFile, stdoutPipe)

jobStderrLogFile, err := os.Create("<job-id>-err")
// handle err
defer jobStderrLogFile.Close()
_, err = io.Copy(jobStderrLogFile, stderrPipe)
```

The Job Id will be used to uniquely identify a job. It will be generated as a UUIDV4 string in the following way:

```
package main

import (
    "fmt"
    "github.com/google/uuid"
)

// This generates a UUIDV4 string
// https://github.com/google/uuid/blob/bfb86fa49a73e4194d93bea18d7acfe3694438ce/version4.go#L13
func main() {
    id := uuid.New()
    fmt.Println(id.String())
}
```


<strong>PS: All code examples should be interpreted as pseudocode. They are not meant to represent the actual code.</strong>

### Trade-offs

* No network isolation between processes. This might lead to some problems such as network ports outage (i.e. only one process can listen on a port at a given time);
* No resource pagination on the API services. This is not scalable but is not necessary for the first version;
* JWT tokens are mocked;
* No CRUD for users;
* No security regarding malicious executables.

### Future work

* Implement mTLS;
* Use network namespaces;
* Improve authorization roles;
* Purge mechanism for obsolete log files;
* Implement login API.

## References

[1] https://crosp.net/blog/software-architecture/clean-architecture-part-2-the-clean-architecture/

[2] https://docs.docker.com/engine/reference/commandline/login/

[3] https://www.techopedia.com/definition/23914/credential-store

[4] https://en.wikipedia.org/wiki/Basic_access_authentication 

[5] https://golang.org/pkg/os/exec/

[6] https://medium.com/bluecore-engineering/implementing-role-based-security-in-a-web-app-89b66d1410e4

[7] https://www.thesslstore.com/blog/tls-1-3-everything-possibly-needed-know/

[8] https://github.com/denji/golang-tls

[9] https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0