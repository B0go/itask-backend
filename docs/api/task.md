# iTask backend API Docs - task

## Create a task

This is the API request used to create a new task.

### Specification

Method: `HTTP PUT`

Path: `/tasks/{uuid}`

Expected request body:
```json
    {
        "title": "task's title",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description",
        "completed" : false
    }
```

Response body:
```json
    {
        "title": "task's title",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description",
        "completed" : false
    }
```

Response codes:

`201 Created` in case of success

`500 Internal Server Error` in case of any error

## Create a list of tasks

This is the API request used to create more then one task in the same time

### Specification

Method: `HTTP POST`

Path: `/tasks`

Expected request body:
```json
    {
        "title": "task's title1",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description1",
        "completed" : true
    },
        {
        "title": "task's title2",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description2",
        "completed" : false
    }
```

Expected response:

```json
    {
        "title": "task's title1",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description1",
        "completed" : true
    },
        {
        "title": "task's title2",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description2",
        "completed" : false
    }
```

Response codes:

`201 Created` in case of success

`500 Internal Server Error` in case of any error

## Delete a task

This is the API request used to delete a task

### Specification

Method: `HTTP DELETE`

Path: `/tasks/{uuid}`

No request body is expected for this request
No response body is expected for this request

Response codes:

`200 OK` in case of success

`500 Internal Server Error` in case of any error

## Get a specific task

This is the API request used to get data of a specific task

### Specification

Method: `HTTP GET`

Path: `/tasks/{uuid}`

No request body is expected for this request

Expected response body:
```json
    {
        "title": "task's title",
        "creation": "2022-07-09",
        "dueDate": "2022-08-09",
        "description": "description",
        "completed" : false
    }
```

Response codes:

`200 OK` in case of success

`500 Internal Server Error` in case of any error

## Get a group of tasks based on query params

This is the API request used to get a group of tasks based on query params

Method: `HTTP GET`

Path: `/tasks`

Quesy Params: `?limit=5&oldest=2020-01-01`

No request body is expected for this request

Expected response body:
```json
    {
        "title": "task's title1",
        "creation": "2020-01-09",
        "dueDate": "2022-08-09",
        "description": "description1",
        "completed" : true
    },
        {
        "title": "task's title2",
        "creation": "2021-03-09",
        "dueDate": "2022-08-09",
        "description": "description2",
        "completed" : false
    }
``` 

Response codes:

`200 OK` in case of success

`500 Internal Server Error` in case of any error
