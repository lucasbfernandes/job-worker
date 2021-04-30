# Create Job

### Name

<strong>job-worker exec</strong> - Create a new job execution.

### Synopsis

<strong>job-worker exec -s SERVER_URL -u USERNAME -p PASSWORD -c EXECUTABLE [ARG...]</strong>

### Description

* <strong>-s:</strong> Server url. Must start with `https`, otherwise command will return with an error.
* <strong>-u:</strong> Username;
* <strong>-p:</strong> Password;
* <strong>-c:</strong> Command that will be executed. Starts with a linux executable name and can be followed by an array of arguments.

<strong>PS:</strong> The CLI will forward the exact same input it received from the user to the `command` JSON field that will be sent to the Server (This is for the sake of simplicity).
Examples:

```
    Command: job-worker exec -s https://server-url.com -u user -p pass -c bin/sh -c "echo hello world"
    
    Will become
    
    JSON object: { command: ["bin/sh", "-c", "echo hello world"] }
```

```
    Command: job-worker exec -s https://server-url.com -u user -p pass -c bin/sh -c ls -la
    
    Will become
    
    JSON object: { command: ["bin/sh", "-c", "ls", "-la"] }
```

### Examples

<strong>Command:</strong> `job-worker exec -s https://server-url.com -u user -p pass -c bin/sh -c ls -la`

<strong>Expected successful output:</strong>
```
    Job created successfuly. Id: 1d655e68-aae0-43d2-adc7-b47c81f1b37e.
```

<strong>Expected authentication error output:</strong>
```
    Failed to create job.
    Error: Username and/or password are wrong.
```

<strong>Expected server error output:</strong>
```
    Failed to create job.
    Error: <error-message>.
```

## References

[1] https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html