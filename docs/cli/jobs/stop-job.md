# Stop Job

### Name

<strong>job-worker stop</strong> - Stop a Job.

### Synopsis

<strong>job-worker stop -s SERVER_URL -i JOB_ID</strong>

### Description

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.
* <strong>-i:</strong> UUIDV4 string representing the Job ID. This id can be found after running the [list jobs command](list-jobs.md).

### Examples

<strong>Command:</strong> `job-worker stop -s https://server-url.com -i 1d655e68-aae0-43d2-adc7-b47c81f1b37e`

<strong>Expected successful output:</strong>
```
    username:
    password:
    
    Job stopped successfuly.
```

<strong>Expected authentication error output:</strong>
```
    username:
    password:
    
    Failed to stop job.
    Error: Username and/or password are wrong.
```

<strong>Expected server error output:</strong>
```
    username:
    password:
    
    Failed to stop job.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html