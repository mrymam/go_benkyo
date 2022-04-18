# MVCの簡単なAPIを作る

## setup
### install migration tool 
```bash
$ brew install dbmate
```

## build & run dev

```bash
$ make dev
```

## migrate

```bash
// init migrate .env file
$ cp db/migrations/.env.default db/migrations/.env
$ make migrate
```

## API
JSON API

```bash
// Create User
$ curl -X POST -v -H "Content-Type: application/json" -d '{"username":"hogehoge", "password": "hogehogehogehogehogehogehoge"}' localhost:3000/users

// Get User List
$ curl -v localhost:3000/users

// Get a User
$ curl -v localhost:3000/users/1
```