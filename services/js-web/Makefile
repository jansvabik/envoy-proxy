.PHONY: dockerbuild dockerrun dockerstop dockerrm

dockerbuild:
	docker build -f Dockerfile -t js-website-lite:1.0 .

dockerrun:
	docker run -d -ti -p 3000:80 --name js-website-lite js-website-lite:1.0

dockerstop:
	docker stop js-website-lite

dockerrm: dockerstop
	docker container rm js-website-lite
