build:
	docker-compose build
	npm run build --prefix targetted-client/
run:
	docker-compose up -d
stopb:
	docker stop $(docker ps -a -q --filter="name=targetted-back") #TODO
stop:
	docker-compose stop
ps:
	docker ps
r:
	make stop
	make build
	make run
	make ps