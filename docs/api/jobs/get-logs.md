# GET jobs/:id/logs

### Request:
```
Body:
""

Headers:
{
  Authorization: Basic ZGVtbzpwQDU1dzByZA==
}
```

#### Query parameters:

<strong>id:</strong> Job id. This is the uuidv4 id returned from the "POST jobs/" request.

### Success response:
```
Body:
{
  logs: "Process exited with code 1"
}

Status code: 200 Ok
```

#### Body parameters:

<strong>logs:</strong> Job logs. Will be retrieved from stdout and stored as a single string for this initial version.
