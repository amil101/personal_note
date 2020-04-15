# Go-Postgres

This project is simple Personal Note application built in golang and using PostgreSQL as DB.

## Pre-requisite
1. Install golang v1.1 or above.
2. Install PostgreSQL 
4. Code Editor (I recommend to use VS Code with Go extension by Microsoft installed)
5. Postman for calling APIs

## Login Credentials
Use basic authentication with following commands.
* username : appId1    password : appPassword    
* username : appId2    password : appPassword

## PostgreSQL Table
Use .env and Setup Db Connection Configurations.
Make sure Database is running properly.

## How to run
1. Setup $GOPATH environment variable
2. checkout this repo somewhere into your $GOPATH
3. run newly generated executable file (The above command will create a binary executable file)
api will be served at
localhost:[port]/api/v1/ ( default port : 8000)
## Api Requests 
# APIS

## API postman samples

https://www.getpostman.com/collections/f28e01333e357bcfcd9a

## add a note

* api

```node
localhost:8000/api/v1/note
```

* sample body

```node
{
	"content":"sample note"
}
```

* sample response body

```node
{
    "message": "success",
    "note": {
        "ID": 1,
        "CreatedAt": "2020-04-16T03:55:55.2269115+05:30",
        "UpdatedAt": "2020-04-16T03:55:55.2269115+05:30",
        "DeletedAt": null,
        "content": "sample note",
        "archive": false,
        "user_id": "appId1"
    },
    "status": true
}
```

## update a note

* api

```node
localhost:8000/api/v1/note/1
```

* sample body

```node
{
	"content":"sample note updated"
}
```

* sample success response body

```node
{
    "message": "success",
    "status": true
}
```

## get all archived notes of a user


* api

```node
localhost:8000/api/v1/note/archived
```


* sample success response body

```node
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2020-04-16T03:59:27.445071+05:30",
            "UpdatedAt": "2020-04-16T03:59:27.445071+05:30",
            "DeletedAt": null,
            "content": "sample note updated",
            "archive": true,
            "user_id": "appId1"
        },
        {
            "ID": 5,
            "CreatedAt": "2020-04-16T04:19:50.812094+05:30",
            "UpdatedAt": "2020-04-16T04:19:50.812094+05:30",
            "DeletedAt": null,
            "content": "sample note",
            "archive": true,
            "user_id": "appId1"
        }
    ],
    "message": "success",
    "status": true
}
```

## get all un-archived notes of a user


* api

```node
localhost:8000/api/v1/note/unarchived
```


* sample success response body

```node
{
    "data": [
        {
            "ID": 3,
            "CreatedAt": "2020-04-16T04:19:04.083994+05:30",
            "UpdatedAt": "2020-04-16T04:19:04.083994+05:30",
            "DeletedAt": null,
            "content": "sample note",
            "archive": false,
            "user_id": "appId1"
        },
        {
            "ID": 4,
            "CreatedAt": "2020-04-16T04:19:08.973049+05:30",
            "UpdatedAt": "2020-04-16T04:19:08.973049+05:30",
            "DeletedAt": null,
            "content": "sample note",
            "archive": false,
            "user_id": "appId1"
        }
    ],
    "message": "success",
    "status": true
}
```

## delete a note

### if archived, delete archived note too

* api

```node
localhost:8000/api/v1/note/2
```

* sample success response body

```node
{
    "message": "success",
    "status": true
}
```



## archive a note


* api

```node
localhost:8000/api/v1/note/1/archive
```

* sample success response body

```node
{
    "message": "success",
    "status": true
}
```


## un-archive a note

* api

```node
localhost:8000/api/v1/note/1/unarchive
```

* sample body

* sample success response body

```node
{
    "message": "success",
    "status": true
}
```


# choice of technology and the reasons

* golang
scalability and performance 
easily deployable, memory safe native binary

* postgre
object-relational database system that uses and extends the SQL language combined with many features 
safely store and scale the data workloads

# improvements

* Maintain Good Security Practices

Better to have a good authentication layer like JWT. Here I used basic authentication only.

* Scalability

Need to add pagination, filtering, sorting as well as access large scale notes.

* Improve response and log

Here we have basic response and logs.

* Improve Error handling

Here we usa a few HTTP status and codes.

