.SHELL := /bin/bash

include db/migrations/.env

dev:
	docker compose up
mysql:
	docker compose up sample-mysql
migrate:
	dbmate -u $(DATABASE_URL) --no-dump-schema up
migrate-down:
	dbmate -u $(DATABASE_URL) --no-dump-schema down