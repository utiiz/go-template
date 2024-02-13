db:
	@docker run --name pg-container -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
