# Job Worker Design Document

This document describes the key design choices and trade-offs for the development of the Job Worker, an application
that enables the execution of arbitrary linux processes remotely. <strong>Users that interact with the system are able to start, stop,
query status and also read the stdout logs</strong> of any requested process execution.

## Architecture overview

The application is composed of 2 components: The CLI and the Server. The former is responsible for parsing user input and translating it
into REST HTTPS requests. The Server receives such requests, maintains internal in-memory state and manages the pool of linux processes that are spawned
along with their statuses.

It's possible to see numbers 1, 2 and 3 in the image below, each representing a different step of a user request. Each step of the request will be described in
details later in this document, but for now, here is a summary of what's happening:

* <strong>Step 1</strong>: After logging in using a basic username/password authentication mechanism, the user types a CLI command requesting one of the possible 
  interactions with a linux process.

* <strong>Step 2</strong>: The CLI application parses the user's command and translates it into an HTTPS request for the Server. The API module of the Server receives 
  the request and check if the input is valid and if the user is authorized to perform the requested operation.

* <strong>Step 3</strong>: If all validations pass, the API module will forward the request to the corresponding internal handler (i.e. [interactor](https://crosp.net/blog/software-architecture/clean-architecture-part-2-the-clean-architecture/)) that will 
  do business logic, access the in-memory state and invoke the OS Process API with the Worker Library when necessary.

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

### Security

#### Authentication

#### Authorization

### Managing Linux Processes

### Trade-offs

### Future work