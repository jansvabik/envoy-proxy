.PHONY: dockerbuild dockerrun dockerrm

dockerbuild:
	docker build -f Dockerfile -t go-redir:1.0 .

dockerrun:
	docker run -d -ti -p 9001:80 --name go-redir go-redir:1.0

dockerstop:
	docker stop go-redir

dockerrm: dockerstop
	docker container rm go-redir
