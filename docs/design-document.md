# Job Worker Design Document

This document describes the key design choices and trade-offs for the development of the Job Worker, an application
that enables the execution of arbitrary linux processes remotely. <strong>Users that interact with the system are able to start, stop,
query status and also read the stdout logs</strong> of any requested process execution.

## Architecture overview

The application is composed of 2 components: The CLI and the Server. The former is responsible for parsing user input and translating it
into REST HTTP calls for the latter. The Server on the other hand is responsible for receiving such requests, maintaining state and managing
the pool of linux processes that will be spawned and their statuses.

![Architecture](../assets/images/architecture.png)

## CLI

### User interface

## API / Internals

### Endpoints

### Entities

### Tables

## Worker library

## Trade-offs