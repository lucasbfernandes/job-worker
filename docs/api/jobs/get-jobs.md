# GET jobs/

### Request:
```
Headers:
{
  Authorization: Basic ZGVtbzpwQDU1dzByZA==
}
```

### Success response:
```
Body:
{
  jobs: [
    {
      id: "da7f970f-1a09-4eb6-bfa3-4975105cb4bf",
      command: ["/bin/bash", "-c", "echo hello"],
      author: "username",
      status: "RUNNING",
      createdAt: ""2021-02-03 10:08:02"
    },
    {
      id: "7f75b775-fd91-40d5-8f0f-e61fb797e46f",
      command: ["/bin/bash", "-c", "echo hello"],
      author: "username",
      status: "STOPPED",
      createdAt: ""2021-02-01 09:07:10"
    },
    {
      id: "bcad5ae5-166c-4ee9-8aec-a08f2c46e4eb",
      command: ["/bin/bash", "-c", "ls"],
      author: "username",
      status: "FAILED",
      createdAt: ""2021-02-04 19:07:10"
    },
    {
      id: "1dd53ed8-34fb-469f-a7bd-245b958c86fc",
      command: ["/bin/bash", "-c", "ls"],
      author: "username",
      status: "COMPLETED",
      createdAt: ""2021-02-01 14:17:10"
    }
  ]
}

Status code: 200 Ok
```

#### Body parameters:

<strong>jobs:</strong> Array of jobs. Will be returned without pagination and will contain relevant information from every job
requested by the authenticated user.