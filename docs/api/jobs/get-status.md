# GET /jobs/:id/status

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
  status: "RUNNING"
}
```

#### Body parameters:

<strong>status:</strong> Status of the job. Will be one of the following: `RUNNING`, `FAILED`, `STOPPED`, `COMPLETED`.
