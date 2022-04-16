# MVCの簡単なAPIを作る

## setup
### install migration tool 
```
$ brew install dbmate
```

### init .env file
```
$ cp db/migrations/.env.default db/migrations/.env
```

## build & run

```
$ docker compose up
```

## migrate
```
$ make migrate
```