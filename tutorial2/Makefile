.SHELL := /bin/bash

include db/migrations/.env

migrate:
	dbmate -u $(DATABASE_URL) --no-dump-schema up
migrate-down:
	dbmate -u $(DATABASE_URL) --no-dump-schema down