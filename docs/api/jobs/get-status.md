# GET /jobs/:id/status

### Required permissions:
jobs.get

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
  status: "COMPLETED",
  createdAt: "2021-02-03 10:08:02",
  finishedAt: "2021-02-03 10:28:02",
  author: "username",
  exitStatus: 0
}
```

#### Body parameters:

* <strong>status:</strong> Status of the job. Will be one of the following: `RUNNING`, `FAILED`, `STOPPED`, `COMPLETED`;
* <strong>createdAt:</strong> Job creation date;
* <strong>finishedAt:</strong> Date when job went from the `RUNNING` status to one of the following: `FAILED`, `STOPPED` or `COMPLETED`. Will be
  `null` when job status is `RUNNING`;
* <strong>author:</strong> Job execution requester;
* <strong>exitStatus:</strong> Process exit status. Might be null if currently in the `RUNNING` state.

### Error response:

<strong>Condition:</strong> Inexistent job. 

```
Status code: 404 Not found
```

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
