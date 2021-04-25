# Job Worker Design Document

This document describes the key design choices and trade-offs for the development of the Job Worker, an application
that enables the execution of arbitrary linux processes remotely. <strong>Users that interact with the system are able to start, stop,
query status and also read the stdout logs</strong> of any requested process execution.

## Architecture overview

The application is composed of 2 components: The CLI and the Server. The former is responsible for parsing user input and translating it
into REST HTTPS requests. The Server receives such requests, maintains internal in-memory state and manages the pool of linux processes that are spawned
along with their statuses.

In Figure 1 it's possible to see numbers 1, 2 and 3, each representing a different step of a user request. Each step of the request will be described in
details later in this document, but for now, here is a summary of what is happening:

* Step <strong>1</strong>: After logging in using a basic username/password authentication mechanism, the user types a CLI command requesting one of the possible 
  interactions with a linux process. For this example, let's say the user is requesting the execution of a new process.

* Step <strong>2</strong>: The CLI application parses the user's command and translates it into an HTTPS request for the Server. The API module of the Server receives 
  the request and check if the input is valid and if the user is authorized to perform the requested operation.

* Step <strong>3</strong>: If all validations pass, the API module will forward the request to the corresponding internal handler (i.e. interactors) that will 
  do business logic, accessing in-memory state and accessing the OS process API with the worker library.

![Architecture](../assets/images/architecture.png)

## CLI

### User interface

## API / Internals

### Endpoints

### Entities

### Tables

## Worker library

## Trade-offs