
  # API-User
This project implements a service that provides a **REST API** to work with user data.

Before running the project, create **.env** in the root directory and fill it in as in the **.env.example** file.
---

To start the program you can write a command:
```bash
    make dcup
```

To stop the program you can write a command:
```bash
    make dcdown
```
You can view the service and database logs with commands:
```bash
    make logapp
```
and
```bash
    make logdb
```
---
Example JSON object user:
```json
{
    "id" : int,
    "name" : string,
    "surname" : string,
    "gender" : enum gender,
    "status" : enum gender,
    "birth_date" : format "2006-01-02",
    "creat_date" : format "2006-01-02"
}
```
---
## Request format
The API supports POST requests with Content-type equal to `application/json` with the following parameters:
| Parameter       | Description      |
|-----------------|------------------|
| Request format  | application/json |
| Response format | application/json |
| Encoding        | UTF-8            |
---
## Enum
User status:
 - active
 - banned
 - deleted

User gender:
 - active
 - banned
 - deleted

---
# EndPoint:
---
## `/user`
### Methods:
#### GET
Returns a list of all users.

Not required query parameters:

| Query     | Description |
|-----------|----------|
| gender    | enum gender|
| status    | enum status |
| full_name | the hollow name is {name + \" \" + last name} string |
| desc      | sorting in descending order, you can pass a few, the value of the attributes of the user |
| asc       | sorting in descending order, you can pass a few, the value of the attributes of the user |
| limit     | limit the number of outputs, type number |
| offset    | offset, type number |
### POST
Creates a user.

Accepts JSON object:
```json
{
    "name": string,
    "surname": string,
    "gender": enum gender,
    "status": enum status,
    "birth_date": format "2006-01-02"
}
```
---
## `/user/{user_id}`
### Methods:
#### GET
Returns the user by ID.
#### PUT
Change user.
Accepts JSON object:
```json
{
    "name": string,
    "surname": string,
    "gender": enum gender,
    "status": enum status,
    "birth_date": format "2006-01-02"
}
```
#### DELETE
Deletes a user by ID.
