.PHONY: 

POSTGRESQL_URL='postgres://postgres@localhost:5432/itube_users?sslmode=disable'

proto:
	@scripts/protogen.sh
security:
	gosec ./...
migrate:
	migrate -database ${POSTGRESQL_URL} -path migrations up
truncate:
	migrate -database ${POSTGRESQL_URL} -path migrations down
generate: update
	go generate ./...
update:
	go mod vendor