## service prototype with go, created to serve as go practice

uses gin, pq, and jwt

### running

execute
`$ go run ./cmd/api-server`

the postgresql database must be named card2go and the program must be able to connect to it with 'postgresql' as user and 'user' as password, and must have a table named users with columns 'id', 'name', 'password'

### endpoints

GET `/ping`

returns a json object with pong in it

---

POST `/auth`

expects an `application/json` object as a body with `"username"` and `"password"` and returns a token