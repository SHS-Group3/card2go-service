## card2go service thing

### building and running the server
execute
```sh
go run .
```

### running dev server
(requires installing air) execute
```sh
air
```

---

## endpoints


`GET /destinations/(id)`
---
returns one destination
```json
// return body
Destination {
    "id": int,
    "name": String,
    "description": String,
    "address": String
    "is_lodging": boolean,
    "booked": boolean,
    "beds": int,
    "rooms: int,
    "packages": Package {
        "id": int,
        "title": String,
        "description": String,
    },
}
```

`GET /destinations[?page=int]`
---
list of destination feed, returns 20 entries per page

`POST /destinations/(id)/book[/pid]`
---
requires header Authentication to be set, creates a booking for the destination `id`, with optional package selector `pid`
```
Authorization: Bearer <token>
```

`POST /auth`
---
returns a token

request body
```
{
    "username": String,
    "password": String,
}
```
response body
```
{
    "token": String,
}
```

`POST /auth/register`
---
creates a user, returns the body

request body
```
{
    "username": String,
    "password": String,
}
```
response body
```
{
    "id": int,
}
```