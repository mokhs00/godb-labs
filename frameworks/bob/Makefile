MYSQL_ROOT_PASSWORD=root
MYSQL_DATABASE=dblabs
MYSQL_USER=dblabs
MYSQL_PASSWORD=dblabs
MYSQL_PORT=3306
MYSQL_HOST=localhost

# Generate models from database
.PHONY: generate-db-models
generate-db-models:
	sh -c 'MYSQL_DSN=$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp\($(MYSQL_HOST):$(MYSQL_PORT)\)/$(MYSQL_DATABASE) go run github.com/stephenafamo/bob/gen/bobgen-mysql@latest'
