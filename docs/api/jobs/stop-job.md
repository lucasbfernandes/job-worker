# POST /jobs/:id/stop

### Required permissions:
jobs.stop

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
```

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
