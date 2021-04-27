# GET /jobs/:id/logs

### Request:
```
Headers:
Authorization: Basic ZGVtbzpwQDU1dzByZA==
```

#### Query parameters:

<strong>id:</strong> Job id. This is the uuidv4 id returned from the "POST jobs/" request.

### Success response:
```
Status code: 200 Ok

Body:
{
  logs: "Process exited with code 1"
}
```

#### Body parameters:

<strong>logs:</strong> Job logs. Will be retrieved from stdout and stored as a single string for this initial version.
