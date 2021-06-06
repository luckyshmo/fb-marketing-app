build:
	docker-compose build
	npm run build --prefix targetted-client/
run:
	docker-compose up -d
stop:
	docker-compose stop
ps:
	docker ps