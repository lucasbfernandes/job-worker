# POST jobs/

### Request:
```
Body:
{
  commmand: ["/bin/bash", "-c", "echo hello"]
}

Headers:
{
  Authorization: Basic ZGVtbzpwQDU1dzByZA==
}
```

#### Body parameters:

<strong>command:</strong> Array of strings in the form `["executable", "param1", "param2", "param3]` [1]. The first element
of the array will always be considered as the executable. Validations: NotNull and NotEmpty.

### Success response:
```
Body:
{
  id: "bdf951f2-f0d8-4e5f-a0ea-79f103391ec9"
}

Status code: 201 Created
```

#### Body parameters:

<strong>id:</strong> Job id. Will be generated as uuidv4 and must be used to apply further commands to it.

## References

[1] https://docs.docker.com/engine/reference/builder/