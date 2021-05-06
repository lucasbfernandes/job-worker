# List Jobs

### Name

<strong>job-worker list</strong> - List all Jobs.

### Synopsis

<strong>job-worker list -s SERVER_URL -t API_TOKEN</strong>

### Description:

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.
* <strong>-t:</strong> API token the user possess. This token is a random string with 20 digits, uppercase and lowercase letters and digits;

### Examples

<strong>Command:</strong> `job-worker list -s https://server-url.com -t 6q6Tz5NBELFo5E9iOSEo`

<strong>Expected successful output:</strong>
```
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
    Failed to fetch jobs.
    Error: Invalid username.
```

<strong>Expected server error output:</strong>
```
    Failed to fetch jobs.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html