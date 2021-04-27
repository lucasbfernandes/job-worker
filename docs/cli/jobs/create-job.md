# Create Job

### Name

<strong>job-worker exec</strong> - Create a new job execution.

### Synopsis

<strong>job-worker exec COMMAND [ARG...]</strong>

### Description

* <strong>COMMAND:</strong> Linux executable name that can be followed by an array of arguments.

### Examples

```
    job-worker exec bin/sh -c ls -la

    job-worker exec touch /tmp/example_file

    job-worker exec pwd

    job-worker exec ps -a
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html