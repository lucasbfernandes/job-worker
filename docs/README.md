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

* <strong>Step 1</strong>: After logging in using a basic username/password authentication mechanism, the user types a CLI command requesting one of the possible
  interactions with a linux process.

* <strong>Step 2</strong>: The CLI application parses the user's command and translates it into an HTTPS request for the Server. The API module of the Server receives
  the request, checks if the input is valid and verifies if the user is authorized to perform the requested operation.

* <strong>Step 3</strong>: If all validations pass, the API module will forward the request to the corresponding internal handler [1].
  It will do business logic, access the persistence layer and invoke the OS Process API with the Worker Library when necessary.

![Architecture](../assets/images/architecture.png)

PS: The image depicts a situation where the CLI and the Server are running on separate machines/OS, which is not always the case. This was intended to highlight the necessity
of creating a secure channel between them.

The next sections will describe in details each of key aspects of the 2 components of the system. The design decisions and trade-offs will be shown in a way that by the end of the
document the reader will have a "grey-box" understanding of the entire system (i.e. Some internal nuances will be mentioned, but not all of them for the sake of readability. These
will be further discussed during application development.)

## CLI

The Job Worker CLI is the user interface of the Job Worker Service. Its responsibility is to authenticate the user with the Server, parse commands, generate authenticated HTTPS requests
and exhibit the responses in a structured manner. This section will:

1. Describe which commands are available for the user and how to use them;
2. Describe how the CLI will manage user secrets (i.e. username/password);
4. Show the trade-offs and what could and should be done for future work.

### Commands

### Managing User Secrets

### Trade-offs

### Future work

## Server

The Job Worker Server is responsible for receiving HTTPS requests, applying validations, and executing the requested action if possible. This section will:

1. Describe which actions are available for requests and how the Server expects to receive them;
2. Describe how the Server will handle security concerns (Authentication and Authorization);
3. Explain briefly how the Server will keep track/state of every spawned process;
4. Show the trade-offs and what could and should be done for future work.

### REST API

#### Jobs:
* [Create Job](api/jobs/create-job.md): `POST jobs`
* [Get Jobs](api/jobs/get-jobs.md): `GET jobs`
* [Stop Job](api/jobs/stop-job.md): `POST jobs/:id/stop`
* [Get Job Status](api/jobs/get-status.md): `GET jobs/:id/status`
* [Get Job Logs](api/jobs/get-logs.md): `GET jobs/:id/logs`

#### Login:
* [Login](api/login/login.md): `POST login/`

### Security

The Job Worker Service will use HTTPS (TLS 1.2) + Basic Authentication (i.e. username/password) in its initial version. Authorization
will take form with the following roles: `Admin`, `Maintainer`, `Developer` and `Reader`.

#### Authentication
Every request made to the Server must contain an authorization header in the form `Authorization: Basic <credentials>`,
where `<credentials>` is the base64 encoding of username and password joined by a single colon `:` [3].

After receiving the request, Server will check if the username/password pair is valid and if it matches any given user in
its in-memory database. If the request is not valid, a `401 Unauthorized` will be returned.

#### Authorization

The Job Worker Service will handle authorization using a simple RBAC mechanism. There will be 4 types of roles, each with
a set of permissions associated with it.

<strong>PS: Stub users will be created as seed data when the Server starts. There will be no users CRUD in the initial version. Some permissions won't be enforced until then. </strong>


|  Roles               |  Permissions         | Description |
| :-------------------:| :-------------------:| :-----------: |
|  Admin | jobs.create, jobs.get, jobs.logs, jobs.stop, users.create, users.update | Can update users + all Maintainer permissions|
|  Maintainer | jobs.create, jobs.get, jobs.logs, jobs.stop, users.create | Can create new users + all Developer permissions|
|  Developer | jobs.create, jobs.get, jobs.logs, jobs.stop | Can create/stop jobs, view all jobs and query their logs|
|  Reader | jobs.get | Can view all jobs and their status|

When receiving an authenticated request, the Server will always check if the user has enough permissions to access the resource. For now, users will have only one 
role.

### Managing Linux Processes

### Trade-offs

### Future work

## References

[1] https://crosp.net/blog/software-architecture/clean-architecture-part-2-the-clean-architecture/

[2] https://github.com/jamescooke/restapidocs

[3] https://en.wikipedia.org/wiki/Basic_access_authentication