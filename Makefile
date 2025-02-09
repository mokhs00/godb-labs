MYSQL_CONTAINER_NAME=godb-labs
MYSQL_VERSION=8.0
MYSQL_ROOT_PASSWORD=root
MYSQL_DATABASE=dblabs
MYSQL_USER=dblabs
MYSQL_PASSWORD=dblabs
MYSQL_PORT=3306

# run-mysql: Run MySQL container
.PHONY: run-mysql
run-mysql:
	docker run --name $(MYSQL_CONTAINER_NAME) -d \
		-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
		-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
		-e MYSQL_USER=$(MYSQL_USER) \
		-e MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
		-p $(MYSQL_PORT):3306 \
		--restart unless-stopped \
		mysql:$(MYSQL_VERSION) --default-authentication-plugin=mysql_native_password
