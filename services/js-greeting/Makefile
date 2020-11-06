.PHONY: dockerbuild dockerrun dockerstop dockerrm

dockerbuild:
	docker build -f Dockerfile -t js-greeting:1.0 .

dockerrun:
	docker run -d -ti -p 3001:80 --name js-greeting js-greeting:1.0

dockerstop:
	docker stop js-greeting

dockerrm: dockerstop
	docker container rm js-greeting
