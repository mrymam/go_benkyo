# MVCの簡単なAPIを作る

## setup
### install migration tool 
```
$ brew install dbmate
```

## build & run

```
$ docker compose up
```

## migrate

```
// init migrate .env file
$ cp db/migrations/.env.default db/migrations/.env
$ make migrate
```