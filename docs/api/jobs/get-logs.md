# GET /jobs/:id/logs

### Required permissions:
jobs.logs

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

### Error response:

<strong>Condition:</strong> Unauthorized user.

```
Status code: 401 Unauthorized
```

### Error response:

<strong>Condition:</strong> User doesn't have enough permissions.

```
Status code: 403 Forbidden
```

## References

[1] https://jsonapi.org/examples/#error-objects

[2] https://github.com/jamescooke/restapidocs