.PHONY: dockerbuild dockerrun dockerrm

dockerbuild:
	docker build -f Dockerfile -t go-time:1.0 .

dockerrun:
	docker run -d -ti -p 9001:80 --name go-time go-time:1.0

dockerstop:
	docker stop go-time

dockerrm: dockerstop
	docker container rm go-time
