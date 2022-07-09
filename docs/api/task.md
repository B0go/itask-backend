# iTask backend API Docs

## Create a task

This is the API request used to create a new task.

### Specification

Method: `HTTP PUT`

Path: `/tasks/{uuid}`

Expected request body:
```json
    {
        "titulo": "titulo da task",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao",
        "completada" : false
    }
```

Response body:
```json
    {
        "titulo": "titulo da task",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao",
        "completada" : false
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
        "titulo": "titulo da task1",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao1",
        "completada" : true
    },
        {
        "titulo": "titulo da task2",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao2",
        "completada" : false
    }
```

Expected response:

```json
    {
        "titulo": "titulo da task1",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao1",
        "completada" : true
    },
        {
        "titulo": "titulo da task2",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao2",
        "completada" : false
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
        "titulo": "titulo da task",
        "dataCriacao": "2022-07-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao",
        "completada" : false
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
        "titulo": "titulo da task1",
        "dataCriacao": "2020-01-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao1",
        "completada" : true
    },
        {
        "titulo": "titulo da task2",
        "dataCriacao": "2021-03-09",
        "dataPrazo": "2022-08-09",
        "descricao": "descricao2",
        "completada" : false
    }
``` 

Response codes:

`200 OK` in case of success

`500 Internal Server Error` in case of any error
