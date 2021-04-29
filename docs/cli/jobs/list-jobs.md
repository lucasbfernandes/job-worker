# List Jobs

### Name

<strong>job-worker list</strong> - List all Jobs.

### Synopsis

<strong>job-worker list -s SERVER_URL</strong>

### Description:

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.

### Examples

<strong>Command:</strong> `job-worker list -s https://server-url.com`

<strong>Expected successful output:</strong>
```
    username:
    password:

    1:
    id: da7f970f-1a09-4eb6-bfa3-4975105cb4bf,
    command: ["/bin/bash", "-c", "echo hello"]
    author: username
    status: RUNNING
    createdAt: 2021-02-03 10:08:02
    finishedAt: 2021-02-03 10:18:02
    
    2:
    id: 7f75b775-fd91-40d5-8f0f-e61fb797e46f,
    command: ["/bin/bash", "-c", "ls", "la"],
    author: username
    status: STOPPED
    createdAt: 2021-02-01 09:07:10
    finishedAt: 2021-02-01 09:10:10
```

<strong>Expected authentication error output:</strong>
```
    username:
    password:
    
    Failed to fetch jobs.
    Error: Username and/or password are wrong.
```

<strong>Expected server error output:</strong>
```
    username:
    password:
    
    Failed to fetch jobs.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html