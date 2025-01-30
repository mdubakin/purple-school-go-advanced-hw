up:
	@docker-compose -f docker/docker-compose.yml up -d

stop:
	@docker-compose -f docker/docker-compose.yml stop

psql:
	@psql "postgres://user:password@localhost:5432/order?sslmode=disable"
