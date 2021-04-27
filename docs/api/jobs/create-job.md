# POST /jobs

### Required permissions: 
jobs.create

### Request:
```
Headers:
Authorization: Basic ZGVtbzpwQDU1dzByZA==

Body:
{
  commmand: ["/bin/bash", "-c", "echo hello"]
}
```

#### Body parameters:

<strong>command:</strong> Array of strings in the form `["executable", "param1", "param2", "param3]` [1]. The first element
of the array will always be considered as the executable. Validations: NotNull and NotEmpty.

### Success response:
```
Status code: 201 Created

Body:
{
  id: "bdf951f2-f0d8-4e5f-a0ea-79f103391ec9"
}
```

#### Body parameters:

<strong>id:</strong> Job id. Will be generated as uuidv4 and must be used to apply further commands to it.

### Error response:

<strong>Condition:</strong> Invalid command parameter (Null or empty array).

```
Status code: 400 Bad request

Body:
{
  errors: [
    {
      status: "400",
      title:  "Invalid Command",
      detail: "Parameter command must not be null or empty"
    }
  ]
}
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

[1] https://docs.docker.com/engine/reference/builder/#run

[2] https://jsonapi.org/examples/#error-objects

[3] https://github.com/jamescooke/restapidocs