# POST login/

### Request:
```
Body:
{
  username: "user1",
  password: "123456
}
```

#### Body parameters:

* <strong>username:</strong> Login credential. Validations: NotNull and NotEmpty.

* <strong>password:</strong> Login credential. Validations: NotNull and NotEmpty.

### Success response:
```
Body:
""

Status code: 200 Ok
```

### Error response:
```
Body:
""

Status code: 401 Unauthorized
```