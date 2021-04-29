# Get Logs

### Name

<strong>job-worker logs</strong> - Get output logs from a specific Job.

### Synopsis

<strong>job-worker logs -s SERVER_URL -i JOB_ID</strong>

### Description

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.
* <strong>-i:</strong> UUIDV4 string representing the Job ID. This id can be found after running the [list jobs command](list-jobs.md).

### Examples

<strong>Command:</strong> `job-worker logs -s https://server-url.com -i 1d655e68-aae0-43d2-adc7-b47c81f1b37e`

<strong>Expected successful output:</strong>
```
    username:
    password:

    -rw-r--r--  1 lucas.fernandes  staff  0 Apr 29 02:20 a
    -rw-r--r--  1 lucas.fernandes  staff  0 Apr 29 02:20 b
    -rw-r--r--  1 lucas.fernandes  staff  0 Apr 29 02:20 c
    -rw-r--r--  1 lucas.fernandes  staff  0 Apr 29 02:20 file-d
    -rw-r--r--  1 lucas.fernandes  staff  0 Apr 29 02:20 some-file
```

<strong>Expected authentication error output:</strong>
```
    username:
    password:
    
    Failed to fetch logs.
    Error: Username and/or password are wrong.
```

<strong>Expected server error output:</strong>
```
    username:
    password:
    
    Failed to fetch logs.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html