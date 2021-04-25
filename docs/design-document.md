# Job Worker Design Document

This document describes the key design choices and trade-offs for the development of the Job Worker, an application
that enables the execution of arbitrary linux processes remotely. <strong>Users that interact with the system are able to start, stop,
query status and also read the stdout logs</strong> of any requested process execution.

## Architecture overview

The application is composed of 2 components: The CLI and the Server. The former is responsible for parsing user input and translating it
into REST HTTP requests. The Server receives such requests, maintains internal in-memory state and manages the pool of linux processes that are spawned
along with their statuses.

In Figure 1 it's possible to see numbers 1, 2 and 3, each representing a different step of a user request that will be explained in details later in the document:

* Step <span style="color:#512FC9">1</span>:

* Step <span style="color:#512FC9">2</span>:

* Step <span style="color:#512FC9">3</span>:

![Architecture](../assets/images/architecture.png)

## CLI

### User interface

## API / Internals

### Endpoints

### Entities

### Tables

## Worker library

## Trade-offs