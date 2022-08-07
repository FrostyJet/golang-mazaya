migrate-up:
	export $(cat env/.env | xargs) &> /dev/null && sql-migrate up -config=dbconfig.yml -env=development

migrate-down:
	export $(cat env/.env | xargs) &> /dev/null && sql-migrate down -config=dbconfig.yml -env=development