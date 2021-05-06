# Stop Job

### Name

<strong>job-worker stop</strong> - Stop a Job.

### Synopsis

<strong>job-worker stop -s SERVER_URL -t API_TOKEN -i JOB_ID</strong>

### Description

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.
* <strong>-t:</strong> API token the user possess. This token is a random string with 20 digits, uppercase and lowercase letters and digits;
* <strong>-i:</strong> UUIDV4 string representing the Job ID. This id can be found after running the [list jobs command](list-jobs.md).

### Examples

<strong>Command:</strong> `job-worker stop -s https://server-url.com -t 6q6Tz5NBELFo5E9iOSEo -i 1d655e68-aae0-43d2-adc7-b47c81f1b37e`

<strong>Expected successful output:</strong>
```
    Job stopped successfuly.
```

<strong>Expected authentication error output:</strong>
```
    Failed to stop job.
    Error: Invalid username.
```

<strong>Expected server error output:</strong>
```
    Failed to stop job.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html