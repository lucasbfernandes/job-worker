# Job Worker Design Document

This document describes the key design choices and trade-offs for the development of the Job Worker, an application
that enables the execution of arbitrary linux processes remotely. <strong>Users that interact with the system are able to start, stop,
query status and also read the stdout logs</strong> of any requested process execution.

## Architecture overview

The application is composed of 2 components: The CLI and the Server. The former is responsible for parsing user input and translating it
into REST HTTPS requests. The Server receives such requests, maintains internal in-memory state and manages the pool of linux processes that are spawned
along with their statuses.

In Figure 1 it's possible to see numbers 1, 2 and 3, each representing a different step of a user request that will be explained in details later in the document:

* Step <strong>1</strong>: After logging in using a basic username/password authentication mechanism, the user requests the execution of a linux process using the
  CLI command.

* Step <strong>2</strong>: The CLI application parses the user's command and translates it into an HTTPS request for the Server. The API module of the Server receives
the request and check if the user is allowed to perform the requested operation.

* Step <strong>3</strong>: If the user is allowed to perform the operation, the API module will pass the request to an internal handler that will persist it and ask
the worker library to call the OS process API to fulfill the user request.

![Architecture](../assets/images/architecture.png)

## CLI

### User interface

## API / Internals

### Endpoints

### Entities

### Tables

## Worker library

## Trade-offs