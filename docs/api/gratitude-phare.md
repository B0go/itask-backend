# iTask backend API Docs - gratitude phrase

## Create a gratitude phrase

This is the API request used to create a new gratitude phrase.

### Specification

Method: `HTTP PUT`

Path: `/gratitude-phrases/{uuid}`

Expected request body:
```json
    {
        "phrase": "my gratitude phrase",
        "creation": "2022-07-09",
    }
```

Response body:
```json
    {
        "phrase": "my gratitude phrase",
        "creation": "2022-07-09",
    }
```

Response codes:

`201 Created` in case of success

`500 Internal Server Error` in case of any error

## Get a group of gratitude phrases based on query params

This is the API request used to get a group of gratitude phrases based on query params

Method: `HTTP GET`

Path: `/gratitude-phrases`

Quesy Params: `?limit=5&oldest=2020-01-01`

No request body is expected for this request

Expected response body:
```json
    {
        "phrase": "my gratitude phrase1",
        "creation": "2022-07-09",
    },    {
        "phrase": "my gratitude phrase2",
        "creation": "2022-07-09",
    }
``` 

Response codes:

`200 OK` in case of success

`500 Internal Server Error` in case of any error
