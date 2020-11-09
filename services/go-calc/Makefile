.PHONY: dockerbuild dockerrun dockerrm

dockerbuild:
	docker build -f Dockerfile -t go-calc:1.0 .

dockerrun:
	docker run -d -ti -p 9002:80 --name go-calc go-calc:1.0

dockerstop:
	docker stop go-calc

dockerrm: dockerstop
	docker container rm go-calc
