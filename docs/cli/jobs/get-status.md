# Get Status

### Name

<strong>job-worker status</strong> - Get the status of a specific Job.

### Synopsis

<strong>job-worker status -s SERVER_URL -i JOB_ID</strong>

### Description

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.
* <strong>-i:</strong> UUIDV4 string representing the Job ID. This id can be found after running the [list jobs command](list-jobs.md).

### Examples

<strong>Command:</strong> `job-worker status -s https://server-url.com -i 1d655e68-aae0-43d2-adc7-b47c81f1b37e`

<strong>Expected successful output:</strong>
```
    username:
    password:

    status: COMPLETED
    createdAt: 2021-02-03 10:08:02
    finishedAt: 2021-02-03 10:28:02
    author: username
```

<strong>Expected authentication error output:</strong>
```
    username:
    password:
    
    Failed to fetch status.
    Error: Username and/or password are wrong.
```

<strong>Expected server error output:</strong>
```
    username:
    password:
    
    Failed to fetch status.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html